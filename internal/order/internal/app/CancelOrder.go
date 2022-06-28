package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ozon/internal/pkg/api"
)

func (s *Server) CancelOrder(ctx context.Context, orderReq *pb.Order) (*empty.Empty, error) {
	_, err := s.repo.CancelOrder(ctx, orderReq)
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
