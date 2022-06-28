package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"ozon/internal/payment/internal/client"
	pb "ozon/internal/pkg/api"
)

func (s *Server) GetPayment(ctx context.Context, paymentReq *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	dbPayment, err := s.repo.GetPayment(ctx, paymentReq)
	if err != nil {
		return &pb.PaymentResponse{}, err
	}

	pbOrder := &pb.Order{ID: dbPayment.OrderID}
	orderResponse, err := client.GetOrder(ctx, pbOrder)
	if err != nil {
		return &pb.PaymentResponse{}, err
	}

	payment := &pb.Payment{
		ID:    dbPayment.ID,
		UUID:  dbPayment.UUID,
		Order: orderResponse.Order,
		TotalPrice: &pb.TotalPrice{
			Price:    orderResponse.Order.TotalPrice.Price,
			Currency: pb.Currency(orderResponse.Order.TotalPrice.Currency),
		},
		Status:    pb.PaymentStatus(dbPayment.Status),
		CreatedAt: &timestamp.Timestamp{Seconds: dbPayment.CreatedAt.Time.Unix()},
		UpdatedAt: &timestamp.Timestamp{Seconds: dbPayment.UpdatedAt.Time.Unix()},
	}
	paymentResponse := &pb.PaymentResponse{
		Payment: payment,
	}

	return paymentResponse, nil
}
