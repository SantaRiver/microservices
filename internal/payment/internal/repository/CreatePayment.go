package repository

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ozon/internal/pkg/api"
)

func (r *repository) CreatePayment(ctx context.Context, paymentReq *pb.Payment) (*empty.Empty, error) {
	const queryCreate = `
		insert into payments (
			order_id,
			status,
			total_price,
			currency
		) values ($1, $2, $3, $4)
	`
	_, err := r.pool.Exec(ctx, queryCreate,
		paymentReq.Order.ID,
		paymentReq.Status,
		paymentReq.TotalPrice.Price,
		paymentReq.TotalPrice.Currency,
	)
	if err != nil {
		return &empty.Empty{}, errors.New("error create payment: " + err.Error())
	}
	return &empty.Empty{}, nil
}
