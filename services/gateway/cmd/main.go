package main

import (
	"log"

	"github.com/dsantaguida/idle-clicker/services/gateway/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//Temp main for now. Plan to expose to web client TBD
	auth_conn, err := grpc.NewClient("localhost:8082",  grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	bank_conn, err := grpc.NewClient("localhost:8081",  grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	_ = service.CreateIdleClient(auth_conn, bank_conn)
}

