package main

import (
	"log"
	"net"

	pb "github.com/dsantaguida/idle-clicker/proto/bank"
	"github.com/dsantaguida/idle-clicker/services/bank/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
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