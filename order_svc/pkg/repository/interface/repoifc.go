package interfaceRepo

import "github.com/14jasimmtp/cityvibe-microservice/order-svc/pkg/models"



type Repo interface {
	CartTotalAmount(userid uint) (float64, error)
	CartFinalPrice(userid uint) (float64, error)
	DisplayCart(userid uint) ([]models.Cart, error)
	ViewAddress(id uint) ([]models.AddressRes, error)
	CheckAddressExist(userid uint, address uint) bool
	CheckPaymentMethodExist(paymentid uint) bool
	CheckCartExist(userID uint) bool
	OrderFromCart(addressid uint, paymentid, userid uint, price float64) (int, error)
	AddOrderProducts(userID uint, orderid int, cart []models.Cart) error 
	UpdateCartAndStockAfterOrder(userID uint, productID int, quantity float64) error
	GetUserById(id int) (*models.UserDetailsResponse, error)
	UpdateShipmentAndPaymentByOrderID(orderStatus string, paymentStatus string, orderID int) (models.OrderDetails, error)
	CheckOrder(orderid string, userID uint) error
	CancelOrder(orderid, pid string, userID uint) error
	UpdateStock(pid int, quantity int) error
	ReturnAmountToWallet(userID uint, orderID, pid string) error
	UpdateOrderFinalPrice(orderID int, amount float64) error
	UpdateWallet(wallet float64, userID uint) error 
	CancelOrderDetails(userID uint, orderID, pid string) (models.CancelDetails, error)
	GetOrderDetails(userID uint) ([]models.ViewOrderDetails, error)
}
