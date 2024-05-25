package services

import (
	"backend/internal"
	"backend/internal/models"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewAccountsService(t *testing.T) {
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	// Your test code here
	accountsService := NewAccountsService(db)
	assert.NotNil(t, accountsService)

	mock.ExpectationsWereMet()
}

func TestAccountsService_Create(t *testing.T) {
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"accounts\"").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectQuery(
		"SELECT .*",
	).WithArgs("123", 1).WillReturnRows(sqlmock.NewRows([]string{"account_id", "balance"}).AddRow("123", 100))

	accountsService := NewAccountsService(db)

	created, err := accountsService.Create(nil, models.Account{
		AccountID: "123",
		Balance:   100,
	})
	assert.Nil(t, err)
	assert.NotNil(t, created)

	mock.ExpectationsWereMet()
}

func TestAccountsService_CreateMultipleSameFails(t *testing.T) {
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"accounts\"").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectQuery(
		"SELECT .*",
	).WithArgs("123", 1).WillReturnRows(sqlmock.NewRows([]string{"account_id", "balance"}).AddRow("123", 100))

	mock.ExpectBegin()
	mock.ExpectExec("INSERT .*").WillReturnError(errors.New("duplicate key"))
	mock.ExpectRollback()

	accountsService := NewAccountsService(db)

	created, err := accountsService.Create(context.Background(), models.Account{
		AccountID: "123",
		Balance:   100,
	})

	assert.Nil(t, err)
	assert.NotNil(t, created)

	_, err = accountsService.Create(context.Background(), models.Account{
		AccountID: "123",
		Balance:   100,
	})

	assert.NotNil(t, err)
	mock.ExpectationsWereMet()
}

func TestAccountsService_CreateNegativeBalance(t *testing.T) {
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	accountsService := NewAccountsService(db)

	_, err := accountsService.Create(context.Background(), models.Account{
		AccountID: "123",
		Balance:   -100,
	})

	assert.NotNil(t, err)
	mock.ExpectationsWereMet()
}

func TestAccountsService_GetByID(t *testing.T) {
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	mock.ExpectQuery(
		"SELECT .*",
	).WithArgs("123", 1).WillReturnRows(sqlmock.NewRows([]string{"account_id", "balance"}).AddRow("123", 100))

	accountsService := NewAccountsService(db)

	account, err := accountsService.GetByID(context.Background(), "123")
	assert.Nil(t, err)
	assert.NotNil(t, account)

	mock.ExpectationsWereMet()
}

func TestAccountsService_GetByUnknownID(t *testing.T) {
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	mock.ExpectQuery(
		"SELECT .*",
	).WithArgs("123", 1).WillReturnError(gorm.ErrRecordNotFound)

	accountsService := NewAccountsService(db)

	account, err := accountsService.GetByID(context.Background(), "123")
	assert.Nil(t, err)
	assert.Nil(t, account)

	mock.ExpectationsWereMet()
}

func TestAccountsService_SubtractBalance(t *testing.T) {
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	mock.ExpectQuery(
		"SELECT .*",
	).WithArgs("123", 1).WillReturnRows(sqlmock.NewRows([]string{"account_id", "balance"}).AddRow("123", 100))

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"accounts\"").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectQuery(
		"SELECT .*",
	).WithArgs("123", 1).WillReturnRows(sqlmock.NewRows([]string{"account_id", "balance"}).AddRow("123", 50))

	accountsService := NewAccountsService(db)

	account, err := accountsService.SubtractBalance(context.Background(), "123", 50)
	assert.Nil(t, err)
	assert.NotNil(t, account)

	mock.ExpectationsWereMet()
}

func TestAccountsService_SubtractBalanceInsufficientBalance(t *testing.T) {
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	mock.ExpectQuery(
		"SELECT .*",
	).WithArgs("123", 1).WillReturnRows(sqlmock.NewRows([]string{"account_id", "balance"}).AddRow("123", 100))

	accountsService := NewAccountsService(db)

	account, err := accountsService.SubtractBalance(context.Background(), "123", 150)
	assert.NotNil(t, err)
	assert.Nil(t, account)

	mock.ExpectationsWereMet()
}

func TestAccountsService_AddBalance(t *testing.T) {
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	mock.ExpectQuery(
		"SELECT .*",
	).WithArgs("123", 1).WillReturnRows(sqlmock.NewRows([]string{"account_id", "balance"}).AddRow("123", 100))

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"accounts\"").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectQuery(
		"SELECT .*",
	).WithArgs("123", 1).WillReturnRows(sqlmock.NewRows([]string{"account_id", "balance"}).AddRow("123", 150))

	accountsService := NewAccountsService(db)

	account, err := accountsService.AddBalance(context.Background(), "123", 50)
	assert.Nil(t, err)
	assert.NotNil(t, account)

	mock.ExpectationsWereMet()
}
