package core

import (
	"bytes"
	"context"

	"github.com/Github-Aiko/Aiko-Core/common"
	"github.com/Github-Aiko/Aiko-Core/common/net"
	"github.com/Github-Aiko/Aiko-Core/common/net/cnc"
	"github.com/Github-Aiko/Aiko-Core/features/routing"
	"github.com/Github-Aiko/Aiko-Core/transport/internet/udp"
)

// CreateObject creates a new object based on the given Aiko instance and config. The Aiko instance may be nil.
func CreateObject(v *Instance, config interface{}) (interface{}, error) {
	ctx := v.ctx
	if v != nil {
		ctx = toContext(v.ctx, v)
	}
	return common.CreateObject(ctx, config)
}

// StartInstance starts a new Aiko instance with given serialized config.
// By default Aiko only support config in protobuf format, i.e., configFormat = "protobuf". Caller need to load other packages to add JSON support.
//
// Aiko:api:stable
func StartInstance(configFormat string, configBytes []byte) (*Instance, error) {
	config, err := LoadConfig(configFormat, bytes.NewReader(configBytes))
	if err != nil {
		return nil, err
	}
	instance, err := New(config)
	if err != nil {
		return nil, err
	}
	if err := instance.Start(); err != nil {
		return nil, err
	}
	return instance, nil
}

// Dial provides an easy way for upstream caller to create net.Conn through Aiko.
// It dispatches the request to the given destination by the given Aiko instance.
// Since it is under a proxy context, the LocalAddr() and RemoteAddr() in returned net.Conn
// will not show real addresses being used for communication.
//
// Aiko:api:stable
func Dial(ctx context.Context, v *Instance, dest net.Destination) (net.Conn, error) {
	ctx = toContext(ctx, v)

	dispatcher := v.GetFeature(routing.DispatcherType())
	if dispatcher == nil {
		return nil, newError("routing.Dispatcher is not registered in Aiko core")
	}

	r, err := dispatcher.(routing.Dispatcher).Dispatch(ctx, dest)
	if err != nil {
		return nil, err
	}
	var readerOpt cnc.ConnectionOption
	if dest.Network == net.Network_TCP {
		readerOpt = cnc.ConnectionOutputMulti(r.Reader)
	} else {
		readerOpt = cnc.ConnectionOutputMultiUDP(r.Reader)
	}
	return cnc.NewConnection(cnc.ConnectionInputMulti(r.Writer), readerOpt), nil
}

// DialUDP provides a way to exchange UDP packets through Aiko instance to remote servers.
// Since it is under a proxy context, the LocalAddr() in returned PacketConn will not show the real address.
//
// TODO: SetDeadline() / SetReadDeadline() / SetWriteDeadline() are not implemented.
//
// Aiko:api:beta
func DialUDP(ctx context.Context, v *Instance) (net.PacketConn, error) {
	ctx = toContext(ctx, v)

	dispatcher := v.GetFeature(routing.DispatcherType())
	if dispatcher == nil {
		return nil, newError("routing.Dispatcher is not registered in Aiko core")
	}
	return udp.DialDispatcher(ctx, dispatcher.(routing.Dispatcher))
}
