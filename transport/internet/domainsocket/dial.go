//go:build !windows && !wasm
// +build !windows,!wasm

package domainsocket

import (
	"context"

	"github.com/Github-Aiko/Aiko-Core/common"
	"github.com/Github-Aiko/Aiko-Core/common/net"
	"github.com/Github-Aiko/Aiko-Core/transport/internet"
	"github.com/Github-Aiko/Aiko-Core/transport/internet/stat"
	"github.com/Github-Aiko/Aiko-Core/transport/internet/tls"
	"github.com/Github-Aiko/Aiko-Core/transport/internet/xtls"
)

func Dial(ctx context.Context, dest net.Destination, streamSettings *internet.MemoryStreamConfig) (stat.Connection, error) {
	settings := streamSettings.ProtocolSettings.(*Config)
	addr, err := settings.GetUnixAddr()
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		return nil, newError("failed to dial unix: ", settings.Path).Base(err).AtWarning()
	}

	if config := tls.ConfigFromStreamSettings(streamSettings); config != nil {
		return tls.Client(conn, config.GetTLSConfig(tls.WithDestination(dest))), nil
	} else if config := xtls.ConfigFromStreamSettings(streamSettings); config != nil {
		return xtls.Client(conn, config.GetXTLSConfig(xtls.WithDestination(dest))), nil
	}

	return conn, nil
}

func init() {
	common.Must(internet.RegisterTransportDialer(protocolName, Dial))
}
