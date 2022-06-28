package app

import (
	"context"
	pb "ozon/internal/pkg/api"
)

func (s *Server) GetUserList(ctx context.Context, shareReq *pb.User) (*pb.UserListResponse, error) {
	userRes, err := s.repo.GetUserList(ctx, shareReq)
	return userRes, err
}
