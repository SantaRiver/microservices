package app

import (
	"context"
	pb "ozon/internal/pkg/api"
)

func (s *Server) GetUser(ctx context.Context, userReq *pb.User) (userRes *pb.UserResponse, err error) {
	userRes, err = s.repo.GetUser(ctx, userReq)
	return
}
