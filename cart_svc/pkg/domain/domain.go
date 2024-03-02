package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID     uint    `json:"user_id" gorm:"uniquekey; not null"`
	Users      User    `json:"-" gorm:"foreignkey:UserID"`
	ProductID  uint    `json:"product_id"`
	Products   Product `json:"-" gorm:"foreignkey:ProductID"`
	Quantity   float64 `json:"quantity"`
	Price      float64 `json:"price"`
	FinalPrice float64 `json:"final_price" gorm:"default:0"`
}

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