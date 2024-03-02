package interfaceRepo

import "github.com/14jasimmtp/cityvibe-microservice/payment-svc/pkg/models"

type Repo interface {
	GetPaymentDetails(orderID int) (models.Payment, error)
	PaymentAlreadyPaid(orderID int) (bool, error)
	PayMethod(orderID int) (int, error)
	AddRazorPayDetails(orderID int, RazorID string) error
	CheckPaymentStatus(orderID int) (string, error)
	UpdatePaymentDetails(orderID int, paymentID string) error
	CheckVerifiedPayment(orderID int) (bool, error)
	UpdateShipmentAndPaymentByOrderID(orderStatus string, paymentStatus string, orderID int) (models.OrderDetails, error)
}
