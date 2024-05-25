package models

import (
	"github.com/jackc/pgtype"

	"gorm.io/gorm"
)

type Idempotency struct {
	BaseImmutable
	KeyHash             string       `gorm:"not null"`
	HttpResponseCode    int          `gorm:"not null"`
	HttpResponseHeaders pgtype.JSONB `gorm:"type:jsonb;default:'[]';not null"`
	HttpResponseBody    string       `gorm:"type:jsonb;default:'{}';not null"`
}

func (r *Idempotency) AfterFind(tx *gorm.DB) (err error) {
	r.Object = "idempotency"
	return
}
