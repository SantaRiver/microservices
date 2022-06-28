package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ozon/internal/pkg/api"
)

func (s *Server) UpdateProduct(ctx context.Context, shareReq *pb.Product) (*empty.Empty, error) {
	err := s.repo.UpdateProduct(ctx, shareReq)
	return &empty.Empty{}, err
}
