package repository

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ozon/internal/pkg/api"
)

func (r *repository) CreateOrder(ctx context.Context, orderReq *pb.Order) (*empty.Empty, error) {
	const queryCreate = `
		insert into orders (
			user_id,
			basket_id,
			status,
			total_price,
			currency
		) values ($1, $2, $3, $4, $5)
	`
	_, err := r.pool.Exec(ctx, queryCreate,
		orderReq.User.ID,
		orderReq.Basket.ID,
		orderReq.Status,
		orderReq.TotalPrice.Price,
		orderReq.TotalPrice.Currency,
	)
	if err != nil {
		return &empty.Empty{}, errors.New("error create order: " + err.Error())
	}
	return &empty.Empty{}, nil
}
