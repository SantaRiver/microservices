package app

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"ozon/internal/order/internal/client"
	pb "ozon/internal/pkg/api"
)

func (s *Server) CreateOrder(ctx context.Context, orderReq *pb.Order) (*empty.Empty, error) {
	if orderReq.User == nil {
		return nil, errors.New("user is required")
	}
	basketResponse, err := client.GetBasket(ctx, orderReq.User)

	if err != nil {
		return nil, err
	}

	orderReq.User = basketResponse.Basket.User
	orderReq.Basket = basketResponse.Basket
	orderReq.TotalPrice = &pb.TotalPrice{
		Price:    basketResponse.Basket.TotalPrice,
		Currency: 0,
	}

	_, err = s.repo.CreateOrder(ctx, orderReq)
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
