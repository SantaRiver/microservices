package repository

import (
	"context"
	pb "ozon/internal/pkg/api"
	"ozon/internal/user/internal/app"
)

func (r *repository) UpdateUser(ctx context.Context, userReq *pb.User) (err error) {
	const query1 = `
		select id from users where id = $1
	`
	resp, err := r.pool.Query(ctx, query1, userReq.ID)
	if err != nil {
		return
	}
	if !resp.Next() {
		err = app.UserNotFoundErr
		return
	}
	const query2 = `
		update users SET name = $2, surname = $3, patronymic = $4, type = $5, birthday = $6, updated_at = current_timestamp WHERE id = $1
	`
	_, err = r.pool.Exec(ctx, query2,
		userReq.ID,
		userReq.Name,
		userReq.Surname,
		userReq.Patronymic,
		userReq.Type,
		userReq.Birthday.AsTime().Format("2006-01-02"),
	)
	return
}
