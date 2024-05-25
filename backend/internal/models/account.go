package models

import "gorm.io/gorm"

type Account struct {
	BaseSimple
	AccountID string  `gorm:"unique;not null" json:"account_id"`
	Balance   float64 `gorm:"default:0" json:"balance"`
}

func (u *Account) AfterFind(tx *gorm.DB) (err error) {
	u.Object = "account"
	return
}
