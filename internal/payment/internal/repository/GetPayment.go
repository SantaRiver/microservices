package repository

import (
	"context"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

func (r *repository) GetPayment(ctx context.Context, paymentReq *pb.PaymentRequest) (models.DBPayment, error) {
	query := `
		select
			id,
			uuid,
			order_id,
			status,
			total_price,
			currency,
			created_at,
			updated_at
		from payments where id = $1
	`
	row := r.pool.QueryRow(ctx, query, paymentReq.ID)
	var dbOrder models.DBPayment
	err := row.Scan(
		&dbOrder.ID,
		&dbOrder.UUID,
		&dbOrder.OrderID,
		&dbOrder.Status,
		&dbOrder.TotalPrice.Price,
		&dbOrder.TotalPrice.Currency,
		&dbOrder.CreatedAt,
		&dbOrder.UpdatedAt,
	)
	if err != nil {
		return models.DBPayment{}, err
	}
	return dbOrder, err
}
