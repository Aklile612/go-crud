package initalizers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionToDB() {

	dsn := "host=localhost user=postgres password=aklile1996 dbname=go-crud port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
