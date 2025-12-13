package main

import (
	"context"
	"fmt"
	"os"

	"net"

	animalsv1 "github.com/rafikmoreira/animals/pkg/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rs/zerolog/log"
)

type Server struct {
	animalsv1.UnimplementedAnimalsServiceServer
}

func (s *Server) GetAnimal(ctx context.Context, req *animalsv1.GetAnimalRequest) (*animalsv1.GetAnimalResponse, error) {
	log.Info().Msg("Listando animais!")
	return nil, status.Error(codes.Unimplemented, "method GetAnimal not implemented")
}
func (s *Server) CreateAnimal(ctx context.Context, req *animalsv1.CreateAnimalRequest) (*animalsv1.CreateAnimalResponse, error) {
	log.Info().Str("animal_name", req.GetName()).Str("animal_type", req.GetAnimalType().String()).Msg("Dados do animal")
	return nil, status.Error(codes.Unimplemented, "method CreateAnimal not implemented")
}

func main() {
	grpcServer := grpc.NewServer()

	server := new(Server)

	animalsv1.RegisterAnimalsServiceServer(grpcServer, server)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Error().Err(err)
	}

	log.Info().Str("port", port).Msg("Inicializando servidor")
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Error().Err(err)
	}
}
