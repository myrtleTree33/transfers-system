package models

import (
	"time"

	"github.com/segmentio/ksuid"
)

// Base contains common columns for all tables.
type Base struct {
	Object string `gorm:"-" json:"object"`

	// adding swaggerignore:"true" because it caused issue upon
	// generating Open API spec due to the type is in ksuid.KSUID
	ID        ksuid.KSUID `gorm:"type:char(27); not null; primary_key;" json:"id" swaggerignore:"true"`
	CreatedAt time.Time   `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time   `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time  `sql:"index" json:"deleted_at"`
}
