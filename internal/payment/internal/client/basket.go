package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"ozon/config"
	pb "ozon/internal/pkg/api"
)

func GetBasket(ctx context.Context, user *pb.User) (basketResponse *pb.BasketResponse, err error) {
	var addr = "localhost" + config.BasketGRPCServerPort
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Basket Client connection error: %v", err)
		}
	}(conn)
	UserClient := pb.NewBasketServiceClient(conn)
	basketResponse, err = UserClient.GetBasket(ctx, user)
	if err != nil {
		return &pb.BasketResponse{}, err
	}
	return
}
