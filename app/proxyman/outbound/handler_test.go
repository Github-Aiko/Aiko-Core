package outbound_test

import (
	"context"
	"testing"

	"github.com/Github-Aiko/Aiko-Core/app/policy"
	. "github.com/Github-Aiko/Aiko-Core/app/proxyman/outbound"
	"github.com/Github-Aiko/Aiko-Core/app/stats"
	"github.com/Github-Aiko/Aiko-Core/common/net"
	"github.com/Github-Aiko/Aiko-Core/common/serial"
	core "github.com/Github-Aiko/Aiko-Core/core"
	"github.com/Github-Aiko/Aiko-Core/features/outbound"
	"github.com/Github-Aiko/Aiko-Core/proxy/freedom"
	"github.com/Github-Aiko/Aiko-Core/transport/internet/stat"
)

func TestInterfaces(t *testing.T) {
	_ = (outbound.Handler)(new(Handler))
	_ = (outbound.Manager)(new(Manager))
}

const xrayKey core.XrayKey = 1

func TestOutboundWithoutStatCounter(t *testing.T) {
	config := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&stats.Config{}),
			serial.ToTypedMessage(&policy.Config{
				System: &policy.SystemPolicy{
					Stats: &policy.SystemPolicy_Stats{
						InboundUplink: true,
					},
				},
			}),
		},
	}

	v, _ := core.New(config)
	v.AddFeature((outbound.Manager)(new(Manager)))
	ctx := context.WithValue(context.Background(), xrayKey, v)
	h, _ := NewHandler(ctx, &core.OutboundHandlerConfig{
		Tag:           "tag",
		ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
	})
	conn, _ := h.(*Handler).Dial(ctx, net.TCPDestination(net.DomainAddress("localhost"), 13146))
	_, ok := conn.(*stat.CounterConnection)
	if ok {
		t.Errorf("Expected conn to not be CounterConnection")
	}
}

func TestOutboundWithStatCounter(t *testing.T) {
	config := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&stats.Config{}),
			serial.ToTypedMessage(&policy.Config{
				System: &policy.SystemPolicy{
					Stats: &policy.SystemPolicy_Stats{
						OutboundUplink:   true,
						OutboundDownlink: true,
					},
				},
			}),
		},
	}

	v, _ := core.New(config)
	v.AddFeature((outbound.Manager)(new(Manager)))
	ctx := context.WithValue(context.Background(), xrayKey, v)
	h, _ := NewHandler(ctx, &core.OutboundHandlerConfig{
		Tag:           "tag",
		ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
	})
	conn, _ := h.(*Handler).Dial(ctx, net.TCPDestination(net.DomainAddress("localhost"), 13146))
	_, ok := conn.(*stat.CounterConnection)
	if !ok {
		t.Errorf("Expected conn to be CounterConnection")
	}
}
