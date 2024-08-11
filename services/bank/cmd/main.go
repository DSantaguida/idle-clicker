package main

import (
	"log"
	"net"

	"github.com/dsantaguida/idle-clicker/pkg/config"
	pb "github.com/dsantaguida/idle-clicker/proto/bank"
	"github.com/dsantaguida/idle-clicker/services/bank/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := config.GetConfig("./services/bank/config/")
	if err != nil {
		log.Fatal("Failed to get config: ", err)
		return
	}

	log.Printf("Creating listener on port: %s", config.Server.Port)

	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		log.Fatalln("Failed to create listener:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterBankServer(s, &service.BankServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}