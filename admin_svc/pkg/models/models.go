package models

type Admin struct {
	ID          uint   `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6,max=20"`
	TokenString string `json:"token"`
}

type AdminLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type AdminOrder struct {
	UserID    int `json:"user_id" validate:"required,number"`
	OrderID   int `json:"order_id" validate:"required,number"`
	ProductID int `json:"product_id" validate:"required,number"`
}

type TimePeriod struct {
	Year      string
	Month     string
	Week      string
	Startdate string
	EndDate   string
}

type DashBoardUser struct {
	TotalUsers  int   `json:"Totaluser"`
	BlockedUser []int32 `json:"Blocked users"`
}
type DashBoardProduct struct {
	TotalProducts       int   `json:"Totalproduct"`
	OutofStockProductID []int32 `json:"Outofstock products"`
	LowStockProductsID  []int32 `json:"less Stock Products"`
}
type DashboardOrder struct {
	DeliveredOrderProducts int
	PendingOrderProducts   int
	CancelledOrderProducts int
	TotalOrderItems        int
	TotalOrderQuantity     int
}
type DashboardRevenue struct {
	TodayRevenue float64
	MonthRevenue float64
	YearRevenue  float64
}
type DashboardAmount struct {
	CreditedAmount float64
	PendingAmount  float64
}
type CompleteAdminDashboard struct {
	DashboardUser    DashBoardUser
	DashboardProduct DashBoardProduct
	DashboardOrder   DashboardOrder
	DashboardRevenue DashboardRevenue
	DashboardAmount  DashboardAmount
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

type ViewAdminOrderDetails struct {
	OrderDetails        AdminOrderDetails
	OrderProductDetails []OrderProductDetails
}

type AdminOrderDetails struct {
	UserID        int
	Id            string
	FinalPrice    float64
	PaymentMethod string
	PaymentStatus string
}

type OrderProductDetails struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	OrderStatus string  `json:"order_status"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}