package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "ozon/internal/pkg/api"
)

func (s *Server) DeleteUser(ctx context.Context, userReq *pb.User) (*empty.Empty, error) {
	err := s.repo.DeleteUser(ctx, userReq)
	return &emptypb.Empty{}, err
}
