package service

import (
	"context"

	bankService "github.com/dsantaguida/idle-clicker/proto/bank"
	"github.com/dsantaguida/idle-clicker/services/bank/internal/models"
)

func (b *BankServiceServer) CreateBank(ctx context.Context, bankRequest *bankService.BankRequest) (*bankService.BankResponse, error) {
	bank := models.CreateBank(bankRequest.Bank.Id, int(bankRequest.Bank.Value))

	bank, err := b.db.CreateBankEntry(ctx, bank)
	if err != nil {
		return nil, err
	}

	return &bankService.BankResponse{Bank: b.bankModelToProto(bank)}, nil
}

func (b *BankServiceServer) GetBankData(ctx context.Context, bankRequest *bankService.GetBankDataRequest) (*bankService.BankResponse, error) {
	bank := models.CreateBank(bankRequest.Id, 0)

	bank, err := b.db.FindBankById(ctx, bank.Id)
	if err != nil {
		return nil, err
	}

	return &bankService.BankResponse{Bank: b.bankModelToProto(bank)}, nil
}

func (b *BankServiceServer) SetBankData(ctx context.Context, bankRequest *bankService.BankRequest) (*bankService.BankResponse, error) {
	bank := models.CreateBank(bankRequest.Bank.Id, int(bankRequest.Bank.Value))

	newBank, err := b.db.UpdateBank(ctx, bank)
	if err != nil {
		return nil, err
	}

	return &bankService.BankResponse{Bank: b.bankModelToProto(newBank)}, nil
}

func (*BankServiceServer) bankModelToProto(bank *models.Bank) (*bankService.Bank){
	return &bankService.Bank{Id: bank.Id, Value: int64(bank.Value)}
} 