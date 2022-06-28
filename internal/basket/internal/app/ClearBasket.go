package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ozon/internal/pkg/api"
)

func (s *Server) ClearBasket(ctx context.Context, user *pb.User) (*empty.Empty, error) {
	err := s.repo.ClearBasket(ctx, user)
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
