package repository

import (
	"errors"
	"fmt"

	"github.com/14jasimmtp/cityvibe-microservice/payment-svc/pkg/models"
	interfaceRepo "github.com/14jasimmtp/cityvibe-microservice/payment-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) interfaceRepo.Repo{
	return &Repo{DB: db}
}

func (r *Repo) GetPaymentDetails(orderID int) (models.Payment, error) {
	var Paymentdt models.Payment
	query := r.DB.Raw(`SELECT users.firstname,orders.total_price as final_price,users.phone FROM orders INNER JOIN users ON orders.user_id=users.id WHERE orders.id = ? `, orderID).Scan(&Paymentdt)
	if query.Error != nil {
		return models.Payment{}, errors.New(`something went wrong`)
	}
	if query.RowsAffected < 1 {
		fmt.Println("hi")
		return models.Payment{}, errors.New(`no orders foun with this id `)
	}
	return Paymentdt, nil
}

func (r *Repo) PaymentAlreadyPaid(orderID int) (bool, error) {
	var paymentStatus string
	query := r.DB.Raw(`SELECT payment_status from orders where id = ? `, orderID).Scan(&paymentStatus)
	if query.Error != nil {
		return false, errors.New(`something went wrong`)
	}
	if query.RowsAffected < 1 {
		return false, errors.New(`no orders found with this id `)
	}
	if paymentStatus == `paid` {
		return true, nil
	}
	return false, nil
}

func (r *Repo) PayMethod(orderID int) (int, error) {
	var id int
	query := r.DB.Raw(`SELECT payment_method_id FROM orders WHERE id = ?`, orderID).Scan(&id)
	if query.Error != nil {
		return 0, errors.New(`something went wrong`)
	}
	if query.RowsAffected < 1 {
		return 0, errors.New(`no orders foun with this id `)
	}
	return id, nil
}

func (r *Repo) AddRazorPayDetails(orderID int, RazorID string) error {
	err := r.DB.Exec("INSERT INTO razor_pays (order_id,razor_id) VALUES (?,?)", orderID, RazorID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) CheckPaymentStatus(orderID int) (string, error) {
	var paymentStatus string
	err := r.DB.Raw(`SELECT payment_status FROM orders WHERE id = $1`, orderID).Scan(&paymentStatus).Error
	if err != nil {
		return "", err
	}
	return paymentStatus, nil
}
func (r *Repo) UpdatePaymentDetails(orderID int, paymentID string) error {
	err := r.DB.Exec("UPDATE razor_pays set payment_id = ? WHERE order_id= ?", paymentID, orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) CheckVerifiedPayment(orderID int) (bool, error) {
	var payment string
	query := r.DB.Raw(`SELECT payment_status FROM orders WHERE id = ?`, orderID).Scan(&payment)
	if query.Error != nil {
		return false, errors.New(`something went wrong`)
	}
	if query.RowsAffected == 0 {
		return false, errors.New(`no orders found`)
	}
	if payment == "paid" {
		return true, nil
	}

	return false, nil
}

func (r *Repo) UpdateShipmentAndPaymentByOrderID(orderStatus string, paymentStatus string, orderID int) (models.OrderDetails, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var details models.OrderDetails

	err := tx.Raw("UPDATE orders SET payment_status = ? WHERE id = ? RETURNING total_price", paymentStatus, orderID).Scan(&details.FinalPrice).Error
	if err != nil {
		tx.Rollback()
		return models.OrderDetails{}, err
	}

	err = tx.Exec("UPDATE order_items SET order_status = ? WHERE order_id = ?", orderStatus, orderID).Error
	if err != nil {
		tx.Rollback()
		return models.OrderDetails{}, errors.New(`something went wrong`)
	}

	details.Id = orderID
	details.PaymentMethod = "Razorpay"
	details.PaymentStatus = "paid"

	tx.Commit()
	return details, nil
}
