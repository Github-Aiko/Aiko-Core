package core_test

import (
	"testing"

	"github.com/Github-Aiko/Aiko-Core/app/dispatcher"
	"github.com/Github-Aiko/Aiko-Core/app/proxyman"
	"github.com/Github-Aiko/Aiko-Core/common"
	"github.com/Github-Aiko/Aiko-Core/common/net"
	"github.com/Github-Aiko/Aiko-Core/common/protocol"
	"github.com/Github-Aiko/Aiko-Core/common/serial"
	"github.com/Github-Aiko/Aiko-Core/common/uuid"
	. "github.com/Github-Aiko/Aiko-Core/core"
	"github.com/Github-Aiko/Aiko-Core/features/dns"
	"github.com/Github-Aiko/Aiko-Core/features/dns/localdns"
	_ "github.com/Github-Aiko/Aiko-Core/main/distro/all"
	"github.com/Github-Aiko/Aiko-Core/proxy/dokodemo"
	"github.com/Github-Aiko/Aiko-Core/proxy/vmess"
	"github.com/Github-Aiko/Aiko-Core/proxy/vmess/outbound"
	"github.com/Github-Aiko/Aiko-Core/testing/servers/tcp"
	"github.com/golang/protobuf/proto"
)

func TestAikoDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestAikoClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{
						Range: []*net.PortRange{net.SinglePortRange(port)},
					},
					Listen: net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}
