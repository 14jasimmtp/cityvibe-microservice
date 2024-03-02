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

type OrderItem struct {
	ID          uint    `json:"id" gorm:"primaryKey;not null"`
	OrderID     uint    `json:"order_id"`
	Order       Order   `json:"-" gorm:"foreignkey:OrderID;constraint:OnDelete:CASCADE"`
	ProductID   uint    `json:"product_id"`
	Products    Product `json:"-" gorm:"foreignkey:ProductID"`
	UserID      int     `json:"user_id"`
	User        User    `json:"-" gorm:"foreignkey:UserID"`
	OrderStatus string  `json:"order_status" gorm:"default:'pending'"`
	Quantity    float64 `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
	FinalPrice  float64 `json:"final_price"`
}

type PaymentMethod struct {
	ID          uint   `json:"id" gorm:"primarykey;not null"`
	PaymentMode string `json:"payment_mode" gorm:"unique; not null"`
}

type Product struct {
	ID          uint     `json:"id" gorm:"unique;not null"`
	Name        string   `json:"name" form:"name"`
	Description string   `json:"description" form:"description"`
	CategoryID  uint     `json:"category_id" form:"category_id"`
	Category    Category `json:"-" gorm:"foreignkey:CategoryID;"`
	SizeID      int      `json:"size_id" form:"size_id"`
	Size        Size     `json:"-" gorm:"foriegnkey:SizeID;"`
	Stock       int      `json:"stock" form:"stock"`
	Price       float64  `json:"price" form:"price"`
	OfferPrize  float64  `json:"offerprice"`
	Color       string   `json:"color" form:"color"`
	Deleted     bool     `json:"delete" gorm:"default:false"`
	ImageURL    string   `json:"imageurl"`
}

type Category struct {
	ID       uint   `json:"id" gorm:"unique; not null"`
	Category string `json:"category" gorm:"unique; not null"`
}

type Size struct {
	ID   uint   `json:"id" gorm:"unique; not null"`
	Size string `json:"size" gorm:"unique; not null"`
}