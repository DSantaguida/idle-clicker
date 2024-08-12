package test

import (
	"testing"

	"github.com/dsantaguida/idle-clicker/pkg/config"
	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"github.com/dsantaguida/idle-clicker/services/bank/internal/db"
	"github.com/dsantaguida/idle-clicker/services/bank/internal/models"
	"golang.org/x/net/context"
)

func Setup(t *testing.T) (*db.BankRepository, context.Context) {
	c, err := config.GetConfig("../config/", "config_test")
	if err != nil {
		t.Fatal("Failed to find test config: ", err)
	}

	db, err := db.CreateBankRepository(c.Db)
	if err != nil {
		t.Fatal("Failed to create database connection: ", err)
	}

	return db, context.TODO()
}

func TestCreateBankEntry(t *testing.T) {
	db, ctx := Setup(t)
	defer db.Close()

	banks := map[string]int {
		"1": 1,
		"2": 36,
		"3": 72, 
		"4": 0,
		"5": 12,
	}

	//Create banks
	for k, v := range banks {
		bank := &models.Bank{Id: k, Value: v}
		_ = CreateBaseBank(bank, ctx, db, t)
	}

	//Try to create bank that already exists
	bank := &models.Bank{Id: "1", Value: 0}
	_, err := db.CreateBankEntry(ctx, bank)
	if err != idle_errors.ErrBankAlreadyExists {
		t.Fatal("Failed to detect bank that already exists.")
	}
}

func CreateBaseBank(bank *models.Bank, ctx context.Context, db *db.BankRepository, t *testing.T) (*models.Bank) {
	createdBank, err := db.CreateBankEntry(ctx, bank)
	if err != nil {
		t.Fatal(err)
	}
	if createdBank == nil {
		t.Fatal("Failed to create banks")
	}
	return createdBank
}

func TestFindBankEntry(t *testing.T) {
	db, ctx := Setup(t)

	//Create a base bank
	bank := &models.Bank{Id: "101", Value: 1}
	_ = CreateBaseBank(bank, ctx, db, t)


	//Find bank that exists
	foundBank, err := db.FindBankById(ctx, "101")
	if err != nil {
		t.Fatal(err)
	}
	if foundBank == nil {
		t.Fatal("Failed to find bank")
	}

	//Find bank that doesn't exist
	foundBank, err = db.FindBankById(ctx, "102")
	if err != nil && err != idle_errors.ErrBankNotExist {
		t.Fatal(err)
	}
	if foundBank != nil && err == idle_errors.ErrBankAlreadyExists {
		t.Fatal("Found bank that does not exist")
	}
}

 func TestUpdateBankEntry(t *testing.T) {
 	db, ctx := Setup(t)

	//Create a base bank
	bank := &models.Bank{Id: "201", Value: 0}
	_ = CreateBaseBank(bank, ctx, db, t)

	//Update base bank value
	newBank := &models.Bank{Id: "201", Value: 100}
	_, err := db.UpdateBankEntry(ctx, newBank)
	if err != nil {
		t.Fatal(err)
	}

	//Check base bank new value
	newBank, err = db.FindBankById(ctx, "201")
	if err != nil {
		t.Fatal(err)
	}
	if newBank.Value != 100 {
		t.Fatal("Failed to update bank value")
	}
 }