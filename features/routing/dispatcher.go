package routing

import (
	"context"

	"github.com/Github-Aiko/Aiko-Core/common/net"
	"github.com/Github-Aiko/Aiko-Core/features"
	"github.com/Github-Aiko/Aiko-Core/transport"
)

// Dispatcher is a feature that dispatches inbound requests to outbound handlers based on rules.
// Dispatcher is required to be registered in a Aiko instance to make Aiko function properly.
//
// Aiko:api:stable
type Dispatcher interface {
	features.Feature

	// Dispatch returns a Ray for transporting data for the given request.
	Dispatch(ctx context.Context, dest net.Destination) (*transport.Link, error)
	DispatchLink(ctx context.Context, dest net.Destination, link *transport.Link) error
}

// DispatcherType returns the type of Dispatcher interface. Can be used to implement common.HasType.
//
// Aiko:api:stable
func DispatcherType() interface{} {
	return (*Dispatcher)(nil)
}
