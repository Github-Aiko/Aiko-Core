package conf

import (
	"github.com/Github-Aiko/Aiko-Core/proxy/loopback"
	"github.com/golang/protobuf/proto"
)

type LoopbackConfig struct {
	InboundTag string `json:"inboundTag"`
}

func (l LoopbackConfig) Build() (proto.Message, error) {
	return &loopback.Config{InboundTag: l.InboundTag}, nil
}
