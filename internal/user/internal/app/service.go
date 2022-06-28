package app

import (
	pb "ozon/internal/pkg/api"
)

type Server struct {
	repo Repository
	pb.UnimplementedUserServiceServer
}

func New(repo Repository) *Server {
	return &Server{repo: repo}
}
