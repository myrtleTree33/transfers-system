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
	SubtractBalance(c context.Context, accountID string, amount float64) (*models.Account, error)
	AddBalance(c context.Context, accountID string, amount float64) (*models.Account, error)
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

// SubtractBalance subtracts an amount from an account's balance
func (s *AccountsService) SubtractBalance(c context.Context, accountID string, amount float64) (*models.Account, error) {
	account, err := s.GetByID(c, accountID)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.New("account not found")
	}

	newBalance := account.Balance - amount
	if newBalance < 0 {
		return nil, errors.New("insufficient balance")
	}

	updatedAccount, err := s.updateBalance(c, accountID, newBalance)
	if err != nil {
		return nil, err
	}

	return updatedAccount, nil
}

// AddBalance adds an amount to an account's balance
func (s *AccountsService) AddBalance(c context.Context, accountID string, amount float64) (*models.Account, error) {
	account, err := s.GetByID(c, accountID)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.New("account not found")
	}

	newBalance := account.Balance + amount
	if newBalance < 0 {
		return nil, errors.New("insufficient balance")
	}

	updatedAccount, err := s.updateBalance(c, accountID, newBalance)
	if err != nil {
		return nil, err
	}

	return updatedAccount, nil
}

// updateBalance updates an account's balance
func (s *AccountsService) updateBalance(c context.Context, accountID string, balance float64) (*models.Account, error) {
	if err := s.db.Model(&models.Account{}).Where("account_id = ?", accountID).Update("balance", balance).Error; err != nil {
		return nil, err
	}
	updatedAccount, err := s.GetByID(c, accountID)
	if err != nil {
		return nil, err
	}
	return updatedAccount, nil
}
