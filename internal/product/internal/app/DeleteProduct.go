package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "ozon/internal/pkg/api"
)

func (s *Server) DeleteProduct(ctx context.Context, userReq *pb.Product) (*empty.Empty, error) {
	err := s.repo.DeleteProduct(ctx, userReq)
	return &emptypb.Empty{}, err
}
