package udp

import (
	"github.com/Github-Aiko/Aiko-Core/common/buf"
	"github.com/Github-Aiko/Aiko-Core/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
