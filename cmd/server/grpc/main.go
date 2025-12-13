package main

import (
	"context"
	"log"
	"net"

	animalsv1 "github.com/rafikmoreira/animals/pkg/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	animalsv1.UnimplementedAnimalsServiceServer
}

func (s *Server) GetAnimal(ctx context.Context, req *animalsv1.GetAnimalRequest) (*animalsv1.GetAnimalResponse, error) {
	log.Print(req)
	return nil, status.Error(codes.Unimplemented, "method GetAnimal not implemented")
}
func (s *Server) CreateAnimal(ctx context.Context, req *animalsv1.CreateAnimalRequest) (*animalsv1.CreateAnimalResponse, error) {
	log.Print(req)
	return nil, status.Error(codes.Unimplemented, "method CreateAnimal not implemented")
}

func main() {
	grpcServer := grpc.NewServer()

	server := new(Server)

	animalsv1.RegisterAnimalsServiceServer(grpcServer, server)

	port := ":9091"

	listener, err := net.Listen("tcp", port)

	if err != nil {
		panic(err.Error())
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		panic(err.Error())
	}
}
