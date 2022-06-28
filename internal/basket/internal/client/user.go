package client

import (
	"context"
	"google.golang.org/grpc"
	"ozon/config"
	pb "ozon/internal/pkg/api"
)

func GetUser(ctx context.Context, user *pb.User) (userResponse *pb.UserResponse, err error) {
	var addr = "localhost" + config.UserGRPCServerPort
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	UserClient := pb.NewUserServiceClient(conn)
	userResponse, err = UserClient.GetUser(ctx, user)
	if err != nil {
		return
	}
	return
}
