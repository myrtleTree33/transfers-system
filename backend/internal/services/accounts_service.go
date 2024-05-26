package services

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

//go:generate mockery --name=IAccountsService
type IAccountsService interface {
	Create(c context.Context, account models.Account) (*models.Account, error)
	GetByID(c context.Context, accountID string) (*models.Account, error)
	TransferBalance(c context.Context, fromAccountID, toAccountID string, amount float64) error
}

// AccountsService is a service that handles account operations
type AccountsService struct {
	db *gorm.DB
}

// NewAccountsService creates a new AccountsService
func NewAccountsService(db *gorm.DB) IAccountsService {
	return &AccountsService{db: db}
}

// Create creates a new account
func (s *AccountsService) Create(c context.Context, account models.Account) (*models.Account, error) {
	if account.Balance < 0 {
		return nil, errors.New("balance cannot be negative")
	}

	if err := s.db.Create(&account).Error; err != nil {
		return nil, err
	}

	createdAccount, err := s.GetByID(c, account.AccountID)
	if err != nil {
		return nil, err
	}

	return createdAccount, nil
}

// GetByID retrieves an account by its ID
func (s *AccountsService) GetByID(c context.Context, accountID string) (*models.Account, error) {
	account := &models.Account{}
	if err := s.db.Where("account_id = ?", accountID).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return account, nil
}

// TransferBalance transfers an amount from one account to another atomically
func (s *AccountsService) TransferBalance(c context.Context, fromAccountID, toAccountID string, amount float64) error {
	// Transaction ensures that the operation is atomic
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Amount must be positive
		if amount <= 0 {
			return errors.New("amount must be positive")
		}

		// Retrieve the from account
		fromAccount := &models.Account{}
		if err := tx.Where("account_id = ?", fromAccountID).First(&fromAccount).Error; err != nil {
			return err
		}

		// If the from account is not found, return an error
		if fromAccount == nil {
			return errors.New("from account not found")
		}

		// Retrieve the to account
		toAccount := &models.Account{}
		if err := tx.Where("account_id = ?", toAccountID).First(&toAccount).Error; err != nil {
			return err
		}

		// If the to account is not found, return an error
		if toAccount == nil {
			return errors.New("to account not found")
		}

		// If the from account has insufficient balance, return an error
		if fromAccount.Balance < amount {
			return errors.New("insufficient balance")
		}

		// Update the from account's balance
		newFromBalance := fromAccount.Balance - amount
		if err := tx.Model(&models.Account{}).Where("account_id = ?", fromAccountID).Update("balance", newFromBalance).Error; err != nil {
			return err
		}

		// Update the to account's balance
		newToBalance := toAccount.Balance + amount
		if err := tx.Model(&models.Account{}).Where("account_id = ?", toAccountID).Update("balance", newToBalance).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}
