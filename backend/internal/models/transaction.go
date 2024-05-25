package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Base
	SourceAccountID      string         `gorm:"not null" json:"source_account_id,omitempty"`
	DestinationAccountID string         `gorm:"not null" json:"destination_account_id,omitempty"`
	Amount               float64        `gorm:"not null" json:"amount,omitempty"`
	CreatedAt            time.Time      `gorm:"not null;default:now()" json:"created_at,omitempty"`
	UpdatedAt            time.Time      `gorm:"not null;default:now()" json:"updated_at,omitempty"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (u *Transaction) AfterFind(tx *gorm.DB) (err error) {
	u.Object = "transaction"
	return
}
