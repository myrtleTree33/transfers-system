package models

import (
	"time"
)

// Base contains common columns for all tables.
type BaseComposite struct {
	Object string `gorm:"-" json:"object"`

	// adding swaggerignore:"true" because it caused issue upon
	// generating Open API spec due to the type is in ksuid.KSUID
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
