package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

type Repository interface {
	CreateOrder(ctx context.Context, orderReq *pb.Order) (*empty.Empty, error)
	GetOrder(ctx context.Context, orderReq *pb.Order) (models.DBOrder, error)
	CancelOrder(ctx context.Context, orderReq *pb.Order) (*empty.Empty, error)
	ConfirmOrder(ctx context.Context, orderReq *pb.Order) (*empty.Empty, error)
}
