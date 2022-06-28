package app

import (
	"context"
	pb "ozon/internal/pkg/api"
)

func (s *Server) GetProduct(ctx context.Context, userReq *pb.Product) (userRes *pb.ProductResponse, err error) {
	userRes, err = s.repo.GetProduct(ctx, userReq)
	return
}
