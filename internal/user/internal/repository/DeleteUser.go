package repository

import (
	"context"
	pb "ozon/internal/pkg/api"
	"ozon/internal/user/internal/app"
)

func (r *repository) DeleteUser(ctx context.Context, userReq *pb.User) (err error) {
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
		delete from users where id = $1
	`
	_, err = r.pool.Exec(ctx, query2, userReq.ID)

	return
}
