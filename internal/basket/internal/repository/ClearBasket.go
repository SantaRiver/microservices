package repository

import (
	"context"
	"errors"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

func (r *repository) ClearBasket(ctx context.Context, user *pb.User) (err error) {
	queryCheck := `select id, products from baskets where user_id = $1`
	row := r.pool.QueryRow(ctx, queryCheck, user.ID)
	var basket models.DBBasket
	err = row.Scan(&basket.ID, &basket.Products)
	if err != nil {
		return errors.New("basket not found")
	}
	if &basket.Products == nil || basket.Products == "[]" {
		return errors.New("basket already is empty")
	}
	queryClear := `update baskets set products = $2 where user_id = $1`
	_, err = r.pool.Exec(ctx, queryClear, user.ID, "[]")
	if err != nil {
		return
	}
	return
}
