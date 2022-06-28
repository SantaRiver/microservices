package app

import (
	"context"
	"errors"
	pb "ozon/internal/pkg/api"
)

var UserNotFoundErr = errors.New("user not found")

type Repository interface {
	CreateUser(ctx context.Context, user *pb.User) (err error)
	UpdateUser(ctx context.Context, user *pb.User) (err error)
	GetUser(ctx context.Context, user *pb.User) (response *pb.UserResponse, err error)
	GetUserList(ctx context.Context, user *pb.User) (response *pb.UserListResponse, err error)
	DeleteUser(ctx context.Context, user *pb.User) (err error)
}
