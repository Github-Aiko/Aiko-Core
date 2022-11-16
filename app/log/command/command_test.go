package command_test

import (
	"context"
	"testing"

	"github.com/Github-Aiko/Aiko-Core/app/dispatcher"
	"github.com/Github-Aiko/Aiko-Core/app/log"
	. "github.com/Github-Aiko/Aiko-Core/app/log/command"
	"github.com/Github-Aiko/Aiko-Core/app/proxyman"
	_ "github.com/Github-Aiko/Aiko-Core/app/proxyman/inbound"
	_ "github.com/Github-Aiko/Aiko-Core/app/proxyman/outbound"
	"github.com/Github-Aiko/Aiko-Core/common"
	"github.com/Github-Aiko/Aiko-Core/common/serial"
	"github.com/Github-Aiko/Aiko-Core/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
