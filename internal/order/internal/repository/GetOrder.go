package repository

import (
	"context"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

func (r *repository) GetOrder(ctx context.Context, orderReq *pb.Order) (models.DBOrder, error) {
	query := `
		select
			id,
			uuid,
			user_id,
			basket_id,
			status,
			total_price,
			currency,
			error,
			created_at,
			updated_at
		from orders where id = $1
	`
	row := r.pool.QueryRow(ctx, query, orderReq.ID)
	var dbOrder models.DBOrder
	err := row.Scan(
		&dbOrder.ID,
		&dbOrder.UUID,
		&dbOrder.UserID,
		&dbOrder.BasketID,
		&dbOrder.Status,
		&dbOrder.TotalPrice.Price,
		&dbOrder.TotalPrice.Currency,
		&dbOrder.Error,
		&dbOrder.CreatedAt,
		&dbOrder.UpdatedAt,
	)
	if err != nil {
		return models.DBOrder{}, err
	}
	return dbOrder, err
}
