package models

import (
	"time"
)

// BaseSimple contains common columns for all tables, except ID.
type BaseSimple struct {
	Object    string     `gorm:"-" json:"object"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
