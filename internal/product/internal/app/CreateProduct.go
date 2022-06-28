package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ozon/internal/pkg/api"
)

func (s *Server) CreateProduct(ctx context.Context, product *pb.Product) (*empty.Empty, error) {
	err := s.repo.CreateProduct(ctx, product)
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
