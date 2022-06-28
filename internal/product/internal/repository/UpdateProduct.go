package repository

import (
	"context"
	pb "ozon/internal/pkg/api"
	"ozon/internal/product/internal/app"
)

func (r *repository) UpdateProduct(ctx context.Context, productReq *pb.Product) (err error) {
	const query1 = `
		select id from users where id = $1
	`
	resp, err := r.pool.Query(ctx, query1, productReq.ID)
	if err != nil {
		return
	}
	if !resp.Next() {
		err = app.ProductNotFoundErr
		return
	}
	const query2 = `
		update products SET 
		name = $2,
		description = $3,
		price = $4,
		image = $5,
		category = $6,
		subcategory = $7,
		brand 	= $8,
		model = $9,
		color = $10,
		size = $11,
		material = $12,
		country = $13,
		updated_at = current_timestamp 
		WHERE id = $1
	`
	_, err = r.pool.Exec(ctx, query2,
		productReq.ID,
		productReq.Name,
		productReq.Description,
		productReq.Price,
		productReq.Image,
		productReq.Category,
		productReq.Subcategory,
		productReq.Brand,
		productReq.Model,
		productReq.Color,
		productReq.Size,
		productReq.Material,
		productReq.Country,
	)
	return
}
