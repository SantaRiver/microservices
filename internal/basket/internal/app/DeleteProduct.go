package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"ozon/internal/basket/internal/client"
	pb "ozon/internal/pkg/api"
)

func (s *Server) DeleteProduct(ctx context.Context, pr *pb.BasketProductRequest) (*empty.Empty, error) {
	userResponse, err := client.GetUser(ctx, pr.User)
	if err != nil {
		return nil, err
	}
	productResponse, err := client.GetProduct(ctx, pr.Product)
	if err != nil {
		return nil, err
	}
	fullPr := pb.BasketProductRequest{
		User:    userResponse.User,
		Product: productResponse.Product,
	}
	err = s.repo.DeleteProduct(ctx, &fullPr)
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
