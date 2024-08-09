package service

import (
	"context"
	"log"

	bankService "github.com/dsantaguida/idle-clicker/proto/bank"
)

func (BankServer) CreateBank(context.Context, *bankService.BankRequest) (*bankService.CreateBankResponse, error) {
	log.Print("Create Bank")
	
	return &bankService.CreateBankResponse{}, nil
}

func (BankServer) GetBankData(context.Context, *bankService.BankRequest) (*bankService.GetBankDataResponse, error) {
	log.Print("Get Bank Data")

	return &bankService.GetBankDataResponse{}, nil
}

func (BankServer) SetBankData(context.Context, *bankService.SetBankDataRequest) (*bankService.SetBankDataResponse, error) {
	log.Print("Set Bank Data")

	return &bankService.SetBankDataResponse{}, nil
}