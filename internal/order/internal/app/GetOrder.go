package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"ozon/internal/order/internal/client"
	pb "ozon/internal/pkg/api"
)

func (s *Server) GetOrder(ctx context.Context, orderReq *pb.Order) (*pb.OrderResponse, error) {
	dbOrder, err := s.repo.GetOrder(ctx, orderReq)
	if err != nil {
		return &pb.OrderResponse{}, err
	}

	pbUser := &pb.User{ID: dbOrder.UserID}
	basketResponse, err := client.GetBasket(ctx, pbUser)

	order := &pb.Order{
		ID:     dbOrder.ID,
		User:   basketResponse.Basket.User,
		Basket: basketResponse.Basket,
		TotalPrice: &pb.TotalPrice{
			Price:    basketResponse.Basket.TotalPrice,
			Currency: pb.Currency(dbOrder.TotalPrice.Currency),
		},
		Status:       pb.OrderStatus(dbOrder.Status),
		ErrorMessage: dbOrder.Error.String,
		CreatedAt:    &timestamp.Timestamp{Seconds: dbOrder.CreatedAt.Time.Unix()},
		UpdatedAt:    &timestamp.Timestamp{Seconds: dbOrder.UpdatedAt.Time.Unix()},
	}
	orderResponse := &pb.OrderResponse{
		Order: order,
	}

	return orderResponse, nil
}
