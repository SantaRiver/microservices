package repository

import (
	"context"
	pb "ozon/internal/pkg/api"
)

func (r *repository) CreateProduct(ctx context.Context, product *pb.Product) (err error) {
	const query = `
		insert into products (
		"name",
		"description",
		"price",
		"image",
		"category",
		"subcategory",
		"brand",
		"model",
		"color",
		"size",
		"material",
		"country"
		) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
	`
	_, err = r.pool.Exec(ctx, query,
		product.Name,
		product.Description,
		product.Price,
		product.Image,
		product.Category,
		product.Subcategory,
		product.Brand,
		product.Model,
		product.Color,
		product.Size,
		product.Material,
		product.Country,
	)
	return err
}
