package services

import (
	"backend/internal/models"
	"context"
	"errors"
	"fmt"

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
	if transaction.Amount <= 0 {
		return nil, errors.New("amount cannot be negative")
	}

	// Transfer balance atomically
	if err := s.accountsService.TransferBalance(c, transaction.SourceAccountID, transaction.DestinationAccountID, transaction.Amount); err != nil {
		return nil, fmt.Errorf("failed to transfer balance: %w", err)
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
