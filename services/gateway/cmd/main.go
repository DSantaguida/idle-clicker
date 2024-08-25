package main

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/dsantaguida/idle-clicker/pkg/interceptors/logging"
	"github.com/dsantaguida/idle-clicker/services/gateway/internal/api/router"
	"github.com/dsantaguida/idle-clicker/services/gateway/internal/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger := logging.CreateClientLogInterceptor()

	//TODO: Remove hardcoded address and port
	auth_conn, err := grpc.NewClient(
		"localhost:8080",  
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		logger,
	)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	bank_conn, err := grpc.NewClient(
		"localhost:8081",  
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		logger,
	)
	if err != nil {
		log.Fatal().Msg(err.Error())
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

