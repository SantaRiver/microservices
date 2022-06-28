package app

import (
	"context"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

type Repository interface {
	GetBasket(ctx context.Context, pbUser *pb.User) (basket models.DBBasket, err error)
	ClearBasket(ctx context.Context, pbUser *pb.User) (err error)
	AddProduct(ctx context.Context, pr *pb.BasketProductRequest) (err error)
	DeleteProduct(ctx context.Context, pr *pb.BasketProductRequest) (err error)
}
