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
	pb "ozon/internal/pkg/api"
	"ozon/internal/user/internal/app"
	"ozon/internal/user/internal/db"
	"ozon/internal/user/internal/mw"
	"ozon/internal/user/internal/repository"
	"time"
)

func RunRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, config.UserRustServerPort, opts)
	if err != nil {
		panic(err)
	}
	log.Printf("User Rust Server listening at " + config.UserRustServerPort)
	if err := http.ListenAndServe(config.UserRustServerPort, mux); err != nil {
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
	lis, err := net.Listen("tcp", "localhost"+config.UserGRPCServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("User gRPC Server listening at " + config.UserGRPCServerPort)

	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(midware.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
