package db

import (
	"log"

	"github.com/14jasimmtp/cityvibe-microservice/product_svc/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(url string) *gorm.DB {
	db,err:=gorm.Open(postgres.Open(url),&gorm.Config{})
	if err != nil {
		log.Print("error connecting db")
	}

	db.AutoMigrate(&domain.Product{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Size{})

	return db
}
