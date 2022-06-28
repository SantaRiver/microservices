package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"ozon/config"
	pb "ozon/internal/pkg/api"
)

func GetOrder(ctx context.Context, user *pb.Order) (orderResponse *pb.OrderResponse, err error) {
	var addr = "localhost" + config.OrderGRPCServerPort
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Order Client connection error: %v", err)
		}
	}(conn)
	UserClient := pb.NewOrderServiceClient(conn)
	orderResponse, err = UserClient.GetOrder(ctx, user)
	if err != nil {
		return &pb.OrderResponse{}, err
	}
	return
}
