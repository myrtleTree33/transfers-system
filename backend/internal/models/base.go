package models

import (
	"time"

	"github.com/segmentio/ksuid"
)

// Base contains common columns for all tables.
type Base struct {
	Object    string       `gorm:"-" json:"object,omitempty"`
	ID        *ksuid.KSUID `gorm:"type:char(27); not null; primary_key;" json:"id,omitempty" swaggerignore:"true"`
	CreatedAt *time.Time   `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt *time.Time   `gorm:"not null" json:"updated_at,omitempty"`
	DeletedAt *time.Time   `sql:"index" json:"deleted_at,omitempty"`
}
