package repository

import (
	"context"
	"errors"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

func (r *repository) GetBasket(ctx context.Context, pbUser *pb.User) (basket models.DBBasket, err error) {
	const query = `
		select 
		id,
		user_id,
		products,
		created_at,
		updated_at
		from baskets where "user_id" = $1 limit 1
	`
	row := r.pool.QueryRow(ctx, query, pbUser.ID)
	var dbBasket models.DBBasket
	err = row.Scan(
		&dbBasket.ID,
		&dbBasket.UserID,
		&dbBasket.Products,
		&dbBasket.CreatedAt,
		&dbBasket.UpdatedAt,
	)
	if err != nil {
		return models.DBBasket{}, errors.New("basket not found")
	}
	basket = dbBasket
	return
}
