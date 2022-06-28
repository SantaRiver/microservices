package repository

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
	"ozon/internal/user/internal/app"
)

func (r *repository) GetUser(ctx context.Context, userReq *pb.User) (userRes *pb.UserResponse, err error) {
	const query = `
		select id, name, surname, patronymic, type, birthday, created_at, updated_at
		from users where "id" = $1 limit 1
	`
	row := r.pool.QueryRow(ctx, query, userReq.ID)
	var user models.DBUser
	err = row.Scan(&user.ID, &user.Name, &user.Surname, &user.Patronymic, &user.Type, &user.Birthday, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, app.UserNotFoundErr
	}
	userRes = &pb.UserResponse{
		User: &pb.User{
			ID:         user.ID,
			Name:       user.Name,
			Surname:    user.Surname,
			Patronymic: user.Patronymic,
			Type:       pb.UserType(user.Type.Int16),
			Birthday:   &timestamp.Timestamp{Seconds: user.Birthday.Time.Unix()},
		},
	}
	return
}
