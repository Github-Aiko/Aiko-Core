package outbound

//go:generate go run github.com/Github-Aiko/Aiko-Core/common/errors/errorgen

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"hash/crc64"
	"time"

	"github.com/Github-Aiko/Aiko-Core/common"
	"github.com/Github-Aiko/Aiko-Core/common/buf"
	"github.com/Github-Aiko/Aiko-Core/common/net"
	"github.com/Github-Aiko/Aiko-Core/common/platform"
	"github.com/Github-Aiko/Aiko-Core/common/protocol"
	"github.com/Github-Aiko/Aiko-Core/common/retry"
	"github.com/Github-Aiko/Aiko-Core/common/session"
	"github.com/Github-Aiko/Aiko-Core/common/signal"
	"github.com/Github-Aiko/Aiko-Core/common/task"
	"github.com/Github-Aiko/Aiko-Core/common/xudp"
	core "github.com/Github-Aiko/Aiko-Core/core"
	"github.com/Github-Aiko/Aiko-Core/features/policy"
	"github.com/Github-Aiko/Aiko-Core/proxy/vmess"
	"github.com/Github-Aiko/Aiko-Core/proxy/vmess/encoding"
	"github.com/Github-Aiko/Aiko-Core/transport"
	"github.com/Github-Aiko/Aiko-Core/transport/internet"
	"github.com/Github-Aiko/Aiko-Core/transport/internet/stat"
)

// Handler is an outbound connection handler for VMess protocol.
type Handler struct {
	serverList    *protocol.ServerList
	serverPicker  protocol.ServerPicker
	policyManager policy.Manager
	cone          bool
}

// New creates a new VMess outbound handler.
func New(ctx context.Context, config *Config) (*Handler, error) {
	serverList := protocol.NewServerList()
	for _, rec := range config.Receiver {
		s, err := protocol.NewServerSpecFromPB(rec)
		if err != nil {
			return nil, newError("failed to parse server spec").Base(err)
		}
		serverList.AddServer(s)
	}

	v := core.MustFromContext(ctx)
	handler := &Handler{
		serverList:    serverList,
		serverPicker:  protocol.NewRoundRobinServerPicker(serverList),
		policyManager: v.GetFeature(policy.ManagerType()).(policy.Manager),
		cone:          ctx.Value("cone").(bool),
	}

	return handler, nil
}

// Process implements proxy.Outbound.Process().
func (h *Handler) Process(ctx context.Context, link *transport.Link, dialer internet.Dialer) error {
	var rec *protocol.ServerSpec
	var conn stat.Connection

	err := retry.ExponentialBackoff(5, 200).On(func() error {
		rec = h.serverPicker.PickServer()
		rawConn, err := dialer.Dial(ctx, rec.Destination())
		if err != nil {
			return err
		}
		conn = rawConn

		return nil
	})
	if err != nil {
		return newError("failed to find an available destination").Base(err).AtWarning()
	}
	defer conn.Close()

	outbound := session.OutboundFromContext(ctx)
	if outbound == nil || !outbound.Target.IsValid() {
		return newError("target not specified").AtError()
	}

	target := outbound.Target
	newError("tunneling request to ", target, " via ", rec.Destination().NetAddr()).WriteToLog(session.ExportIDToError(ctx))

	command := protocol.RequestCommandTCP
	if target.Network == net.Network_UDP {
		command = protocol.RequestCommandUDP
	}
	if target.Address.Family().IsDomain() && target.Address.Domain() == "v1.mux.cool" {
		command = protocol.RequestCommandMux
	}

	user := rec.PickUser()
	request := &protocol.RequestHeader{
		Version: encoding.Version,
		User:    user,
		Command: command,
		Address: target.Address,
		Port:    target.Port,
		Option:  protocol.RequestOptionChunkStream,
	}

	account := request.User.Account.(*vmess.MemoryAccount)
	request.Security = account.Security

	if request.Security == protocol.SecurityType_AES128_GCM || request.Security == protocol.SecurityType_NONE || request.Security == protocol.SecurityType_CHACHA20_POLY1305 {
		request.Option.Set(protocol.RequestOptionChunkMasking)
	}

	if shouldEnablePadding(request.Security) && request.Option.Has(protocol.RequestOptionChunkMasking) {
		request.Option.Set(protocol.RequestOptionGlobalPadding)
	}

	if request.Security == protocol.SecurityType_ZERO {
		request.Security = protocol.SecurityType_NONE
		request.Option.Clear(protocol.RequestOptionChunkStream)
		request.Option.Clear(protocol.RequestOptionChunkMasking)
	}

	if account.AuthenticatedLengthExperiment {
		request.Option.Set(protocol.RequestOptionAuthenticatedLength)
	}

	input := link.Reader
	output := link.Writer

	isAEAD := false
	if !aeadDisabled && len(account.AlterIDs) == 0 {
		isAEAD = true
	}

	hashkdf := hmac.New(sha256.New, []byte("VMessBF"))
	hashkdf.Write(account.ID.Bytes())

	behaviorSeed := crc64.Checksum(hashkdf.Sum(nil), crc64.MakeTable(crc64.ISO))

	session := encoding.NewClientSession(ctx, isAEAD, protocol.DefaultIDHash, int64(behaviorSeed))
	sessionPolicy := h.policyManager.ForLevel(request.User.Level)

	ctx, cancel := context.WithCancel(ctx)
	timer := signal.CancelAfterInactivity(ctx, cancel, sessionPolicy.Timeouts.ConnectionIdle)

	if request.Command == protocol.RequestCommandUDP && h.cone && request.Port != 53 && request.Port != 443 {
		request.Command = protocol.RequestCommandMux
		request.Address = net.DomainAddress("v1.mux.cool")
		request.Port = net.Port(666)
	}

	requestDone := func() error {
		defer timer.SetTimeout(sessionPolicy.Timeouts.DownlinkOnly)

		writer := buf.NewBufferedWriter(buf.NewWriter(conn))
		if err := session.EncodeRequestHeader(request, writer); err != nil {
			return newError("failed to encode request").Base(err).AtWarning()
		}

		bodyWriter, err := session.EncodeRequestBody(request, writer)
		if err != nil {
			return newError("failed to start encoding").Base(err)
		}
		bodyWriter2 := bodyWriter
		if request.Command == protocol.RequestCommandMux && request.Port == 666 {
			bodyWriter = xudp.NewPacketWriter(bodyWriter, target)
		}
		if err := buf.CopyOnceTimeout(input, bodyWriter, time.Millisecond*100); err != nil && err != buf.ErrNotTimeoutReader && err != buf.ErrReadTimeout {
			return newError("failed to write first payload").Base(err)
		}

		if err := writer.SetBuffered(false); err != nil {
			return err
		}

		if err := buf.Copy(input, bodyWriter, buf.UpdateActivity(timer)); err != nil {
			return err
		}

		if request.Option.Has(protocol.RequestOptionChunkStream) && !account.NoTerminationSignal {
			if err := bodyWriter2.WriteMultiBuffer(buf.MultiBuffer{}); err != nil {
				return err
			}
		}

		return nil
	}

	responseDone := func() error {
		defer timer.SetTimeout(sessionPolicy.Timeouts.UplinkOnly)

		reader := &buf.BufferedReader{Reader: buf.NewReader(conn)}
		header, err := session.DecodeResponseHeader(reader)
		if err != nil {
			return newError("failed to read header").Base(err)
		}
		h.handleCommand(rec.Destination(), header.Command)

		bodyReader, err := session.DecodeResponseBody(request, reader)
		if err != nil {
			return newError("failed to start encoding response").Base(err)
		}
		if request.Command == protocol.RequestCommandMux && request.Port == 666 {
			bodyReader = xudp.NewPacketReader(&buf.BufferedReader{Reader: bodyReader})
		}

		return buf.Copy(bodyReader, output, buf.UpdateActivity(timer))
	}

	responseDonePost := task.OnSuccess(responseDone, task.Close(output))
	if err := task.Run(ctx, requestDone, responseDonePost); err != nil {
		return newError("connection ends").Base(err)
	}

	return nil
}

var (
	enablePadding = false
	aeadDisabled  = false
)

func shouldEnablePadding(s protocol.SecurityType) bool {
	return enablePadding || s == protocol.SecurityType_AES128_GCM || s == protocol.SecurityType_CHACHA20_POLY1305 || s == protocol.SecurityType_AUTO
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return New(ctx, config.(*Config))
	}))

	const defaultFlagValue = "NOT_DEFINED_AT_ALL"

	paddingValue := platform.NewEnvFlag("xray.vmess.padding").GetValue(func() string { return defaultFlagValue })
	if paddingValue != defaultFlagValue {
		enablePadding = true
	}

	isAeadDisabled := platform.NewEnvFlag("xray.vmess.aead.disabled").GetValue(func() string { return defaultFlagValue })
	if isAeadDisabled == "true" {
		aeadDisabled = true
	}
}
