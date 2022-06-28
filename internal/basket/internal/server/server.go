package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"ozon/config"
	"ozon/internal/basket/internal/app"
	"ozon/internal/basket/internal/db"
	midware "ozon/internal/basket/internal/mw"
	"ozon/internal/basket/internal/repository"
	pb "ozon/internal/pkg/api"
	"time"
)

func RunRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterBasketServiceHandlerFromEndpoint(ctx, mux, config.BasketRustServerPort, opts)
	if err != nil {
		panic(err)
	}
	log.Printf("Basket Rust Server listening at " + config.BasketRustServerPort)
	if err := http.ListenAndServe(config.BasketRustServerPort, mux); err != nil {
		panic(err)
	}
}

func RunGrpc() {
	ctx := context.Background()

	adp, err := db.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	rep := repository.New(adp)
	newServer := app.New(rep)
	lis, err := net.Listen("tcp", "localhost"+config.BasketGRPCServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Basket gRPC Server listening at " + config.BasketGRPCServerPort)

	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(midware.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBasketServiceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
