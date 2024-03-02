package db

import (
	"log"

	"github.com/14jasimmtp/cityvibe-microservice/payment-svc/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(url string) *gorm.DB {
	db,err:=gorm.Open(postgres.Open(url),&gorm.Config{})
	if err != nil {
		log.Print("error connecting db")
	}

	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(&domain.PaymentMethod{})
	db.AutoMigrate(&domain.RazorPay{})
	db.AutoMigrate(&domain.Wallet{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Order{})

	return db
}
