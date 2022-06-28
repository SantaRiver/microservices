package app

import (
	pb "ozon/internal/pkg/api"
)

type Server struct {
	repo Repository
	pb.UnimplementedOrderServiceServer
}

func New(repo Repository) *Server {
	return &Server{repo: repo}
}
