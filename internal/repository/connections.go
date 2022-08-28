package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewGormDB returns the instance for *gorm.DB
func NewGormDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn))
}
