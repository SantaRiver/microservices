package app

import (
	"context"
	"errors"
	pb "ozon/internal/pkg/api"
)

var ProductAlreadyExistErr = errors.New("product already exist")
var ProductNotFoundErr = errors.New("product not found")

type Repository interface {
	CreateProduct(ctx context.Context, product *pb.Product) (err error)
	UpdateProduct(ctx context.Context, product *pb.Product) (err error)
	GetProduct(ctx context.Context, product *pb.Product) (response *pb.ProductResponse, err error)
	GetProductList(ctx context.Context, product *pb.Product) (response *pb.ProductListResponse, err error)
	DeleteProduct(ctx context.Context, product *pb.Product) (err error)
}
