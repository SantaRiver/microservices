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
	"ozon/internal/product/internal/app"
	"ozon/internal/product/internal/db"
	midware "ozon/internal/product/internal/mw"
	"ozon/internal/product/internal/repository"
	"time"
)

func RunRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterProductServiceHandlerFromEndpoint(ctx, mux, config.ProductRustServerPort, opts)
	if err != nil {
		panic(err)
	}
	log.Printf("Product Rest Server listening at " + config.ProductRustServerPort)
	if err := http.ListenAndServe(config.ProductRustServerPort, mux); err != nil {
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
	lis, err := net.Listen("tcp", "localhost"+config.ProductGRPCServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Product gRPC server listening at " + config.ProductGRPCServerPort)

	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(midware.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterProductServiceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
