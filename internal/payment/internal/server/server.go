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
	"ozon/internal/payment/internal/app"
	"ozon/internal/payment/internal/db"
	midware "ozon/internal/payment/internal/mw"
	"ozon/internal/payment/internal/repository"
	pb "ozon/internal/pkg/api"
	"time"
)

func RunRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, config.PaymentRustServerPort, opts)
	if err != nil {
		panic(err)
	}
	log.Printf("Payment Rust Server listening at " + config.PaymentRustServerPort)
	if err := http.ListenAndServe(config.PaymentRustServerPort, mux); err != nil {
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
	lis, err := net.Listen("tcp", "localhost"+config.PaymentGRPCServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Payment gRPC Server listening at " + config.PaymentGRPCServerPort)

	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(midware.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPaymentServiceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
