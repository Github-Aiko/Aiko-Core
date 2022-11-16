package udp

import (
	"github.com/Github-Aiko/Aiko-Core/common"
	"github.com/Github-Aiko/Aiko-Core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
