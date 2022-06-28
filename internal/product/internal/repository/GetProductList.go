package repository

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

func (r *repository) GetProductList(ctx context.Context, productReq *pb.Product) (productListRes *pb.ProductListResponse, err error) {
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
		from products
	`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return
	}
	var productList pb.ProductListResponse
	for rows.Next() {
		var dbProduct models.DBProduct
		err = rows.Scan(
			&dbProduct.ID,
			&dbProduct.Name,
			&dbProduct.Description,
			&dbProduct.Price,
			&dbProduct.Image,
			&dbProduct.Category,
			&dbProduct.Subcategory,
			&dbProduct.Brand,
			&dbProduct.Model,
			&dbProduct.Color,
			&dbProduct.Size,
			&dbProduct.Material,
			&dbProduct.Country,
			&dbProduct.CreatedAt,
			&dbProduct.UpdatedAt,
		)
		if err != nil {
			return
		}
		product := &pb.Product{
			ID:          dbProduct.ID,
			Name:        dbProduct.Name,
			Description: dbProduct.Description,
			Price:       dbProduct.Price,
			Image:       dbProduct.Image,
			Category:    dbProduct.Category,
			Subcategory: dbProduct.Subcategory,
			Brand:       dbProduct.Brand,
			Model:       dbProduct.Model,
			Color:       dbProduct.Color,
			Size:        dbProduct.Size,
			Material:    dbProduct.Material,
			Country:     dbProduct.Country,
			CreatedAt:   &timestamp.Timestamp{Seconds: dbProduct.CreatedAt.Time.Unix()},
			UpdatedAt:   &timestamp.Timestamp{Seconds: dbProduct.UpdatedAt.Time.Unix()},
		}
		productList.Products = append(productList.Products, product)
	}

	return &productList, nil
}
