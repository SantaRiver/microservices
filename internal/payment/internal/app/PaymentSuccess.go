package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ozon/internal/pkg/api"
)

func (s *Server) PaymentSuccess(ctx context.Context, paymentStatusReq *pb.PaymentStatusRequest) (*empty.Empty, error) {
	_, err := s.repo.PaymentSuccess(ctx, paymentStatusReq)
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
