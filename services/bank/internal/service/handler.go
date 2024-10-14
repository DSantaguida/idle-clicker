package service

import (
	"context"

	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"github.com/dsantaguida/idle-clicker/pkg/jwt"
	bankService "github.com/dsantaguida/idle-clicker/proto/bank"
	"github.com/dsantaguida/idle-clicker/services/bank/internal/models"
	"google.golang.org/grpc/metadata"
)

func (b *BankServiceServer) CreateBank(ctx context.Context, bankRequest *bankService.BankRequest) (*bankService.BankResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, idle_errors.ErrNoMetadata
	}

	token := md.Get(jwt.TOKEN_KEY)[0]
	if len(token) == 0 {
		return nil, idle_errors.ErrTokenNotInHeader
	}

	id, err := jwt.ParseId(token)
	if err != nil {
		return nil, err
	}

	bank, err := b.db.CreateBankEntry(ctx, id)
	if err != nil {
		return nil, err
	}

	return &bankService.BankResponse{Bank: b.bankModelToProto(bank)}, nil
}

func (b *BankServiceServer) GetBankData(ctx context.Context, bankRequest *bankService.GetBankDataRequest) (*bankService.BankResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, idle_errors.ErrNoMetadata
	}

	token := md.Get(jwt.TOKEN_KEY)[0]
	if len(token) == 0 {
		return nil, idle_errors.ErrTokenNotInHeader
	}

	id, err := jwt.ParseId(token)
	if err != nil {
		return nil, err
	}

	bank := models.CreateBank(id, 0)

	bank, err = b.db.FindBankById(ctx, bank.Id)
	if err != nil {
		return nil, err
	}

	return &bankService.BankResponse{Bank: b.bankModelToProto(bank)}, nil
}

func (b *BankServiceServer) SetBankData(ctx context.Context, bankRequest *bankService.SetBankDataRequest) (*bankService.BankResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, idle_errors.ErrNoMetadata
	}

	token := md.Get(jwt.TOKEN_KEY)[0]
	if len(token) == 0 {
		return nil, idle_errors.ErrTokenNotInHeader
	}

	id, err := jwt.ParseId(token)
	if err != nil {
		return nil, err
	}

	bank := models.CreateBank(id, int(bankRequest.Value))

	newBank, err := b.db.UpdateBankEntry(ctx, bank)
	if err != nil {
		return nil, err
	}

	return &bankService.BankResponse{Bank: b.bankModelToProto(newBank)}, nil
}

func (*BankServiceServer) bankModelToProto(bank *models.Bank) (*bankService.Bank){
	return &bankService.Bank{Id: bank.Id, Value: int64(bank.Value)}
} 