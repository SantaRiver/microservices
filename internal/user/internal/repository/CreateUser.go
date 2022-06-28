package repository

import (
	"context"
	pb "ozon/internal/pkg/api"
)

func (r *repository) CreateUser(ctx context.Context, user *pb.User) (err error) {
	const query = `
		insert into users
		("name", "surname", "patronymic", "type", "birthday") 
		values ($1, $2, $3, $4, $5);
	`
	_, err = r.pool.Exec(ctx, query, user.Name, user.Surname, user.Patronymic, user.Type, user.Birthday.AsTime().Format("2006-01-02"))
	return err
}
