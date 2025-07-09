package initalizers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
func ConnectionToDB() {
	var err error
	dsn := "host=localhost user=postgres password=aklile1996 dbname=go-crud port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
