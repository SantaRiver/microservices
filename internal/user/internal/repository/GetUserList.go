package repository

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"log"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

func (r *repository) GetUserList(ctx context.Context, userReq *pb.User) (userListRes *pb.UserListResponse, err error) {
	const query = `
		select id, name, surname, patronymic, type, birthday, created_at, updated_at
		from users
	`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return
	}
	var userList pb.UserListResponse
	for rows.Next() {
		var dbUser models.DBUser
		err = rows.Scan(&dbUser.ID, &dbUser.Name, &dbUser.Surname, &dbUser.Patronymic, &dbUser.Type, &dbUser.Birthday, &dbUser.CreatedAt, &dbUser.UpdatedAt)
		if err != nil {
			return
		}
		user := &pb.User{
			ID:         dbUser.ID,
			Name:       dbUser.Name,
			Surname:    dbUser.Surname,
			Patronymic: dbUser.Patronymic,
			Type:       pb.UserType(dbUser.Type.Int16),
			Birthday:   &timestamp.Timestamp{Seconds: dbUser.Birthday.Time.Unix()},
		}
		log.Println(user)
		userList.Users = append(userList.Users, user)
	}

	return &userList, nil
}
