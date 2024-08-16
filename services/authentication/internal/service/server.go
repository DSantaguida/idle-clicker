package service

import (
	pb "github.com/dsantaguida/idle-clicker/proto/authentication"
	"github.com/dsantaguida/idle-clicker/services/authentication/internal/db"
)

type AuthenticationServiceServer struct {
	pb.UnimplementedAuthenticationServiceServer
	db *db.AuthenticationRepository;
}

func CreateServer(db *db.AuthenticationRepository) (*AuthenticationServiceServer){
	return &AuthenticationServiceServer{db: db}
}