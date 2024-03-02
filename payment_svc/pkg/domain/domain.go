package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint    `json:"id" gorm:"unique;not null"`
	Firstname  string  `json:"firstname"`
	Lastname   string  `json:"lastname"`
	Email      string  `json:"email" validate:"email"`
	Phone      string  `json:"phone"`
	Password   string  `json:"-" validate:"min=8,max=20"`
	Blocked    bool    `json:"blocked" gorm:"default:false"`
	Wallet     float64 `json:"wallet" gorm:"default:0"`
}

type Address struct {
	Id        int    `json:"id" gorm:"unique;not null"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"-" gorm:"foreignkey:UserID"`
	Name      string `json:"name" validate:"required"`
	Phone     string `json:"phone" validate:"required min=10 max=10"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}

type Wallet struct {
	ID     uint    `json:"id" gorm:"unique;not null"`
	UserID uint    `json:"user_id"`
	User   User    `json:"-" gorm:"foreignkey:UserID"`
	Amount float64 `json:"Balance" gorm:"default:0;not null"`
}

type PaymentMethod struct {
	ID          uint   `json:"id" gorm:"primarykey;not null"`
	PaymentMode string `json:"payment_mode" gorm:"unique; not null"`
}

type RazorPay struct {
	ID        uint   `json:"id" gorm:"primarykey not null"`
	OrderID   string `json:"order_id" `
	Order     Order  `json:"-" gorm:"foreignkey:OrderID"`
	RazorID   string `json:"razor_id"`
	PaymentID string `json:"payment_id"`
}

type Order struct {
	gorm.Model
	UserID          int           `json:"user_id" gorm:"not null"`
	User            User          `json:"-" gorm:"foreignkey:UserID"`
	AddressID       int           `json:"address_id" gorm:"not null"`
	Address         Address       `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID uint          `json:"payment_method_id"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	PaymentStatus   string        `json:"payment_status" gorm:"default:'not paid'"`
	TotalPrice      float64       `json:"total_price"`
	FinalPrice      float64       `json:"final_price"`
	Approval        bool          `json:"approval" gorm:"default:false"`
}