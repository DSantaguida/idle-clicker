package service

import (
	pb "github.com/dsantaguida/idle-clicker/proto/bank"
	"github.com/dsantaguida/idle-clicker/services/bank/internal/db"
)

type BankServiceServer struct {
	pb.UnimplementedBankServiceServer
	db *db.BankRepository;
}

func CreateServer(db *db.BankRepository) (*BankServiceServer){
	return &BankServiceServer{db: db}
}