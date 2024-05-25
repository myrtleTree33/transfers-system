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

type AccountsService struct {
	db *gorm.DB
}

func NewAccountsService(db *gorm.DB) IAccountsService {
	return &AccountsService{db: db}
}

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

func (s *AccountsService) SubtractBalance(c context.Context, accountID string, amount float64) (*models.Account, error) {
	account, err := s.GetByID(c, accountID)
	if err != nil {
		return nil, err
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

func (s *AccountsService) AddBalance(c context.Context, accountID string, amount float64) (*models.Account, error) {
	account, err := s.GetByID(c, accountID)
	if err != nil {
		return nil, err
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
