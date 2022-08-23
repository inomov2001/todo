package main

import (
	"net"
	"user-service/repository"
	"user-service/server"
	"user-service/service"
	"user-service/userpb"

	"google.golang.org/grpc"
)

func main() {
	db := repository.NewDB()
	service := service.New(db)

	RunGRPCServer(service)
}

func RunGRPCServer(svc service.Service) {
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, server.NewGRPCServer(svc))

	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
