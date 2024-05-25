package internal

import (
	"database/sql"
	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupTestDB() (*sql.DB, sqlmock.Sqlmock, *gorm.DB) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	return mockDb, mock, db
}

type Anything struct{}

func (a Anything) Match(v driver.Value) bool {
	return true
}
