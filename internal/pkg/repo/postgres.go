package repo

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDb() (*gorm.DB, error) {
	dsn := "user=gorm dbname=gorm password=gorm sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
