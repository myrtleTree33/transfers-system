package services

import (
	"backend/internal/models"
	"context"
	"errors"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

// ITransactionsService is the interface for the transactions service
type ITransactionsService interface {
	Create(c context.Context, Transaction models.Transaction) (*models.Transaction, error)
	GetByID(c context.Context, transactionID ksuid.KSUID) (*models.Transaction, error)
}

// TransactionsService is the service that handles transaction operations
type TransactionsService struct {
	db              *gorm.DB
	accountsService IAccountsService
}

// NewTransactionsService creates a new TransactionsService
func NewTransactionsService(db *gorm.DB, accountsService IAccountsService) ITransactionsService {
	return &TransactionsService{
		db:              db,
		accountsService: accountsService,
	}
}

// Create creates a new transaction
func (s *TransactionsService) Create(c context.Context, transaction models.Transaction) (*models.Transaction, error) {
	// Subtract balance from source account first
	// If the source account does not have enough balance, return an error
	_, err := s.accountsService.SubtractBalance(c, transaction.SourceAccountID, transaction.Amount)
	if err != nil {
		return nil, err
	}

	_, err = s.accountsService.AddBalance(c, transaction.DestinationAccountID, transaction.Amount)
	if err != nil {
		return nil, err
	}

	if err := s.db.Create(&transaction).Error; err != nil {
		return nil, err
	}

	createdTransaction, err := s.GetByID(c, *transaction.ID)
	if err != nil {
		return nil, err
	}

	return createdTransaction, nil
}

// GetByID retrieves a transaction by its ID
func (s *TransactionsService) GetByID(c context.Context, transactionID ksuid.KSUID) (*models.Transaction, error) {
	transaction := &models.Transaction{}
	if err := s.db.Where("id = ?", transactionID).First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return transaction, nil
}
