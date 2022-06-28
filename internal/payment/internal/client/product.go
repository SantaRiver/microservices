package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"ozon/config"
	pb "ozon/internal/pkg/api"
)

func GetProduct(ctx context.Context, product *pb.Product) (userResponse *pb.ProductResponse, err error) {
	var addr = "localhost" + config.ProductGRPCServerPort
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Product Client connection error: %v", err)
		}
	}(conn)
	UserClient := pb.NewProductServiceClient(conn)
	userResponse, err = UserClient.GetProduct(ctx, product)
	if err != nil {
		return &pb.ProductResponse{}, err
	}
	return
}
