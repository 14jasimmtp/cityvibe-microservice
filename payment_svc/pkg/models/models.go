package models

type Payment struct {
	Total_price float64
	Final_price float64
	Username    string
	Userphone   string
}

type PaymentVerify struct {
	PaymentID string `json:"payment_id" validate:"required"`
	OrderID   string `json:"order_id" validate:"required"`
	Signature string `json:"signature" validate:"required"`
}

type OrderDetails struct {
	Id            int
	FinalPrice    float64
	PaymentMethod string
	PaymentStatus string
}