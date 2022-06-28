package repository

import (
	"context"
	pb "ozon/internal/pkg/api"
	"ozon/internal/product/internal/app"
)

func (r *repository) DeleteProduct(ctx context.Context, productReq *pb.Product) (err error) {
	const query1 = `
		select id from products where id = $1
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
		delete from users where id = $1
	`
	_, err = r.pool.Exec(ctx, query2, productReq.ID)

	return
}
