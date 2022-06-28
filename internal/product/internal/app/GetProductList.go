package app

import (
	"context"
	pb "ozon/internal/pkg/api"
)

func (s *Server) GetProductList(ctx context.Context, shareReq *pb.Product) (*pb.ProductListResponse, error) {
	userRes, err := s.repo.GetProductList(ctx, shareReq)
	return userRes, err
}
