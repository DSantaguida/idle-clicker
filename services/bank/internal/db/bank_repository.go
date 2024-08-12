package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dsantaguida/idle-clicker/pkg/config"
	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"github.com/dsantaguida/idle-clicker/services/bank/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
)

type BankRepository struct {
	db *sql.DB;
}

func (bankRepository *BankRepository) Close(){
	bankRepository.db.Close()
}

func CreateBankRepository(dbConfig config.DBConfig) (*BankRepository, error){
	url := fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.SslMode)

	db, err := sql.Open(dbConfig.Driver, url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil{
		return nil, err
	}

	return &BankRepository{db: db}, nil
}

func (b *BankRepository) CreateBankEntry(ctx context.Context, bank *models.Bank) (*models.Bank, error){
	existingBank, err := b.FindBankById(ctx, bank.Id)
	if err != nil && err != idle_errors.ErrBankNotExist {
		return nil, err
	} else if existingBank != nil {
		return nil, idle_errors.ErrBankAlreadyExists
	}

	createdBank := &models.Bank{}
	query := fmt.Sprintf(createBankQuery, bank.Id, bank.Value)
	err = b.db.QueryRowContext(ctx, query).Scan(&createdBank.Id, &createdBank.Value)
	if err != nil {
		return nil, errors.Wrap(err, "CreateBankEntry")
	}

	return createdBank, nil
}

func (b *BankRepository) FindBankById(ctx context.Context, id string) (*models.Bank, error) {
	query := fmt.Sprintf(findBankQuery, id)

	existingBank := &models.Bank{}
	err := b.db.QueryRowContext(ctx, query).Scan(&existingBank.Id, &existingBank.Value)
	if err != nil {
		return nil, idle_errors.ErrBankNotExist
	}

	return existingBank, nil
}

func (d *BankRepository) UpdateBankEntry(ctx context.Context, bank *models.Bank) (*models.Bank, error) {
	existingBank, err := d.FindBankById(ctx, bank.Id)
	if err != nil {
		return nil, err
	}

	newBank := &models.Bank{}
	query := fmt.Sprintf(updateBankQuery, bank.Value, existingBank.Id)
	err = d.db.QueryRowContext(ctx, query).Scan(&newBank.Id, &newBank.Value)
	if err != nil {
		return nil, err
	}

	return newBank, nil
}