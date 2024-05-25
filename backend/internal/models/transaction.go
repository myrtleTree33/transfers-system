package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	Base
	SourceAccountID      string  `gorm:"not null" json:"source_account_id,omitempty"`
	DestinationAccountID string  `gorm:"not null" json:"destination_account_id,omitempty"`
	Amount               float64 `gorm:"not null" json:"amount,omitempty"`
}

func (u *Transaction) AfterFind(tx *gorm.DB) (err error) {
	u.Object = "transaction"
	return
}
