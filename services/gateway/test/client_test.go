package test

import (
	"context"
	"log"
	"testing"

	"github.com/dsantaguida/idle-clicker/services/gateway/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Setup() (*service.IdleClient){
	auth_conn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	bank_conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	idleClient := service.CreateIdleClient(auth_conn, bank_conn)
	return idleClient
}

func TestRegisterLoginUpdate(t *testing.T) {
	idleClient := Setup()

	username := "dansan858"
	password := "password123"
	ctx := context.TODO()

	err := idleClient.Register(ctx, username, password)
	if err != nil {
		t.Fatal(err)
	}

	token, value, err := idleClient.Login(ctx, username, password)
	if err != nil {
		t.Fatal(err)
	}
	if value != 0  || len(token) == 0 {
		t.Fatal("failed to login to new account")
	}

	err = idleClient.UpdateBankValue(ctx, token, 100)
	if err != nil {
		t.Fatal(err)
	}
}