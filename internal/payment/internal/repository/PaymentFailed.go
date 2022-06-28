package repository

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

func (r *repository) PaymentFailed(ctx context.Context, paymentStatusReq *pb.PaymentStatusRequest) (*empty.Empty, error) {
	queryCheck := `SELECT id FROM payments WHERE id = $1`
	row := r.pool.QueryRow(ctx, queryCheck, paymentStatusReq.ID)
	var payment models.DBPayment
	err := row.Scan(&payment.ID)
	if err != nil {
		return &empty.Empty{}, errors.New("error fail payment: " + err.Error())
	}
	queryUpdate := `UPDATE payments SET status = $1 WHERE id = $2`
	_, err = r.pool.Exec(ctx, queryUpdate, pb.PaymentStatus_FAILED, paymentStatusReq.ID)
	if err != nil {
		return &empty.Empty{}, errors.New("error fail payment: " + err.Error())
	}
	return &empty.Empty{}, nil
}
