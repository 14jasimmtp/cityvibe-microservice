package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(url string) *gorm.DB {
	db,err:=gorm.Open(postgres.Open(url),&gorm.Config{})
	if err != nil {
		log.Print("error connecting db")
	}

	return db
}
