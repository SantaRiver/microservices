package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

type Repository interface {
	CreatePayment(ctx context.Context, paymentReq *pb.Payment) (*empty.Empty, error)
	GetPayment(ctx context.Context, paymentReq *pb.PaymentRequest) (models.DBPayment, error)
	PaymentSuccess(ctx context.Context, paymentStatusReq *pb.PaymentStatusRequest) (*empty.Empty, error)
	PaymentFailed(ctx context.Context, paymentStatusReq *pb.PaymentStatusRequest) (*empty.Empty, error)
}
