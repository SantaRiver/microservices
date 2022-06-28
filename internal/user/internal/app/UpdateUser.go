package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ozon/internal/pkg/api"
)

func (s *Server) UpdateUser(ctx context.Context, shareReq *pb.User) (*empty.Empty, error) {
	err := s.repo.UpdateUser(ctx, shareReq)
	return &empty.Empty{}, err
}
