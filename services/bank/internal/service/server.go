package service

import (
	pb "github.com/dsantaguida/idle-clicker/proto/bank"
)

type BankServer struct {
	pb.UnimplementedBankServer
}
