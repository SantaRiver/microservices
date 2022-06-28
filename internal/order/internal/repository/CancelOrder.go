package repository

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

func (r *repository) CancelOrder(ctx context.Context, orderReq *pb.Order) (*empty.Empty, error) {
	queryCheck := `SELECT id FROM orders WHERE id = $1`
	row := r.pool.QueryRow(ctx, queryCheck, orderReq.ID)
	var order models.DBOrder
	err := row.Scan(&order.ID)
	if err != nil {
		return &empty.Empty{}, errors.New("error cancel order: " + err.Error())
	}
	queryUpdate := `UPDATE orders SET status = $1 WHERE id = $2`
	_, err = r.pool.Exec(ctx, queryUpdate, pb.OrderStatus_CANCELLED, orderReq.ID)
	if err != nil {
		return &empty.Empty{}, errors.New("error cancel order: " + err.Error())
	}
	return &empty.Empty{}, nil
}
