package services

import (
	"backend/internal"
	"backend/internal/models"
	"backend/internal/services/mocks"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	mocklib "github.com/stretchr/testify/mock"
)

func TestNewTransactionsService(t *testing.T) {
	// Write your test cases here
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	// Your test code here
	accountsService := &mocks.IAccountsService{}
	transactionsService := NewTransactionsService(db, accountsService)

	assert.NotNil(t, transactionsService)

	mock.ExpectationsWereMet()
}

func TestTransactionsService_Create(t *testing.T) {
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	accountsService := &mocks.IAccountsService{}
	transactionsService := NewTransactionsService(db, accountsService)

	accountsService.On("TransferBalance", mocklib.Anything, "123", "456", float64(100)).Return(nil, nil)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT .*").
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	id := ksuid.New()

	mock.ExpectQuery("SELECT .*").
		WithArgs(
			&id,
			1,
		).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "source_account_id", "destination_account_id", "amount"}).
				AddRow(id, "123", "456", 100),
		)

	created, err := transactionsService.Create(
		context.Background(), models.Transaction{
			Base: models.Base{
				ID: &id,
			},
			SourceAccountID:      "123",
			DestinationAccountID: "456",
			Amount:               100,
		})
	assert.Nil(t, err)
	assert.NotNil(t, created)

	mock.ExpectationsWereMet()
}

func TestTransactionsService_AccountBalanceNegative(t *testing.T) {
	sqlDb, mock, db := internal.SetupTestDB()
	defer sqlDb.Close()

	accountsService := &mocks.IAccountsService{}
	transactionsService := NewTransactionsService(db, accountsService)

	accountsService.On("TransferBalance", mocklib.Anything, "123", "456", float64(100)).Return(nil, errors.New("negative balance"))

	created, err := transactionsService.Create(
		context.Background(), models.Transaction{
			SourceAccountID:      "123",
			DestinationAccountID: "456",
			Amount:               100,
		})
	assert.NotNil(t, err)
	assert.Nil(t, created)

	mock.ExpectationsWereMet()
}
