package models

type Cart struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	FinalPrice  float64 `json:"-"`
}

type CartResponse struct {
	TotalPrice float64
	Cart       []Cart
}
type Admin struct {
	ID          uint   `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6,max=20"`
	TokenString string `json:"token"`
}
type UserDetailsResponse struct {
	ID        uint    `json:"ID"`
	Email     string  `json:"email"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	Blocked   bool    `json:"blocked"`
	Wallet    float64 `json:"wallet"`
}

type AddressRes struct {
	ID         int    `json:"id"`
	Name       string `json:"name" validate:"required"`
	House_name string `json:"house_name" validate:"required"`
	Phone      string `json:"phone"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state" validate:"required"`
	Pin        string `json:"pin" validate:"required"`
}

type OrderDetails struct {
	Id            int
	FinalPrice    float64
	PaymentMethod string
	PaymentStatus string
}

type ViewOrderDetails struct {
	OrderDetails        OrderDetails
	OrderProductDetails []OrderProductDetails
}

type OrderProductDetails struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	OrderStatus string  `json:"order_status"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}

type CancelDetails struct {
	OrderStatus   string  `json:"order_status"`
	Quantity      int     `json:"quantity"`
	PaymentStatus string  `json:"payment_status"`
	TotalPrice    float64 `json:"total_price"`
	OrderID       int     `json:"order_id"`
	ProductID     int     `json:"product_id"`
}
