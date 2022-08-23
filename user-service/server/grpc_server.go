package server

import (
	"context"
	"user-service/service"
	"user-service/userpb"
)

func NewGRPCServer(svc service.Service) *GRPCServer {
	return &GRPCServer{
		svc: svc,
	}
}

type GRPCServer struct {
	userpb.UnimplementedUserServiceServer
	svc service.Service
}

func (s *GRPCServer) RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (*userpb.User, error) {
	id, err := s.svc.RegisterUser(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	return &userpb.User{
		Id:   int32(id),
		Name: req.GetName(),
	}, nil
}

func (s *GRPCServer) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.User, error) {
	user, err := s.svc.GetUserByID(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &userpb.User{
		Id:   int32(user.ID),
		Name: user.Name,
	}, nil
}
