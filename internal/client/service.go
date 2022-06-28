package client

import pb "ozon/internal/pkg/api"

type Server struct {
	repo Repository
	pb.UnimplementedPaymentServiceServer
}

func New(repo Repository) *Server {
	return &Server{repo: repo}
}
