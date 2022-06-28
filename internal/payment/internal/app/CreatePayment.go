package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"ozon/internal/payment/internal/client"
	pb "ozon/internal/pkg/api"
)

func (s *Server) CreatePayment(ctx context.Context, paymentCreateReq *pb.PaymentCreateRequest) (*empty.Empty, error) {
	orderResponse, err := client.GetOrder(ctx, &pb.Order{ID: paymentCreateReq.OrderID})
	if err != nil {
		return nil, err
	}

	var paymentReq pb.Payment
	paymentReq.Order = orderResponse.Order
	paymentReq.TotalPrice = &pb.TotalPrice{
		Price:    orderResponse.Order.TotalPrice.Price,
		Currency: orderResponse.Order.TotalPrice.Currency,
	}

	_, err = s.repo.CreatePayment(ctx, &paymentReq)
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
