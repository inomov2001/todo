package main

import (
	"net"
	"todo-service/repository"
	"todo-service/server"
	"todo-service/service"
	"todo-service/todopb"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	repo := repository.NewDB()
	svc := service.New(repo)
	todopb.RegisterTodoServiceServer(s, server.NewGRPCServer(svc))

	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
