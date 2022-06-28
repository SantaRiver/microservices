package repository

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"log"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
	"ozon/internal/product/internal/app"
)

func (r *repository) GetProduct(ctx context.Context, userReq *pb.Product) (userRes *pb.ProductResponse, err error) {
	const query = `
		select
		id,
		name,
		description,
		price,
		image,
		category,
		subcategory,
		brand,
		model,
		color,
		size,
		material,
		country,
		created_at,
		updated_at
		from products where "id" = $1 limit 1
	`
	row := r.pool.QueryRow(ctx, query, userReq.ID)
	var product models.DBProduct
	err = row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Image,
		&product.Category,
		&product.Subcategory,
		&product.Brand,
		&product.Model,
		&product.Color,
		&product.Size,
		&product.Material,
		&product.Country,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	log.Println(&product)
	if err != nil {
		return nil, app.ProductNotFoundErr
	}
	userRes = &pb.ProductResponse{
		Product: &pb.Product{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Image:       product.Image,
			Category:    product.Category,
			Subcategory: product.Subcategory,
			Brand:       product.Brand,
			Model:       product.Model,
			Color:       product.Color,
			Size:        product.Size,
			Material:    product.Material,
			Country:     product.Country,
			CreatedAt:   &timestamp.Timestamp{Seconds: product.CreatedAt.Time.Unix()},
			UpdatedAt:   &timestamp.Timestamp{Seconds: product.UpdatedAt.Time.Unix()},
		},
	}
	return
}
