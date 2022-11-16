package extension

import (
	"context"

	"github.com/Github-Aiko/Aiko-Core/features"
	"github.com/golang/protobuf/proto"
)

type Observatory interface {
	features.Feature

	GetObservation(ctx context.Context) (proto.Message, error)
}

func ObservatoryType() interface{} {
	return (*Observatory)(nil)
}
