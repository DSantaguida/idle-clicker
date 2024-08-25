package main

import (
	"net"

	"github.com/dsantaguida/idle-clicker/pkg/config"
	"github.com/dsantaguida/idle-clicker/pkg/interceptors/logging"
	pb "github.com/dsantaguida/idle-clicker/proto/authentication"
	"github.com/dsantaguida/idle-clicker/services/authentication/internal/db"
	"github.com/dsantaguida/idle-clicker/services/authentication/internal/service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := logging.CreateServerLogInterceptor()

	config, err := config.GetConfig("./services/authentication/config/", "config")
	if err != nil {
		log.Fatal().Msgf("Failed to get config: %s", err)
	}

	db, err := db.CreateAuthenticationRepository(config.Db)
	if err != nil {
		log.Fatal().Msgf("Failed to create db repository: %s", err)
	}
	defer db.Close()

	log.Info().Msgf("Creating listener on port: %s", config.Server.Port)

	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		log.Fatal().Msgf("Failed to create listener: %s", err)
	}

	s := grpc.NewServer(logger)
	reflection.Register(s)

	server := service.CreateServer(db)

	pb.RegisterAuthenticationServiceServer(s, server)
	if err := s.Serve(listener); err != nil {
		log.Fatal().Msgf("Failed to serve: %s", err)
	}
}