package grpc

import (
	"context"
	"example/go_server/pkg/proto"
)

type AdderServer struct {
	proto.UnimplementedAdderServer
}

func (s *AdderServer) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	return &proto.AddResponse{Result: int64(req.GetX()) + int64(req.GetY())}, nil
}
