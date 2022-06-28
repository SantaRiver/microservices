package app

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "ozon/internal/pkg/api"
	"time"
)

func (s *Server) CreateUser(ctx context.Context, user *pb.User) (*empty.Empty, error) {
	err := s.repo.CreateUser(ctx, user)
	yearNow := time.Now().Year()
	if user.Birthday.AsTime().Year() > yearNow {
		return &emptypb.Empty{}, errors.New("birthday is in future")
	}
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
