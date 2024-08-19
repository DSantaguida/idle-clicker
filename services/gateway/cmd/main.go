package main

import (
	"log"
	"net/http"

	"github.com/dsantaguida/idle-clicker/services/gateway/internal/api/router"
	"github.com/dsantaguida/idle-clicker/services/gateway/internal/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//TODO: Remove hardcoded address and port
	auth_conn, err := grpc.NewClient("localhost:8080",  grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	bank_conn, err := grpc.NewClient("localhost:8081",  grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := client.CreateIdleClient(auth_conn, bank_conn)
	defer client.Close()

	router := router.NewRouter(*client)
	
	s := &http.Server {
		Addr: ":8083",
		Handler: router,
	}

	s.ListenAndServe()
}

