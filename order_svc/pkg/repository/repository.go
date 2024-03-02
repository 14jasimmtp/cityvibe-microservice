package repository

import (
	"errors"
	"fmt"

	"github.com/14jasimmtp/cityvibe-microservice/order-svc/pkg/models"
	interfaceRepo "github.com/14jasimmtp/cityvibe-microservice/order-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) interfaceRepo.Repo {
	return &Repo{DB: db}
}

func (r *Repo) CartTotalAmount(userid uint) (float64, error) {
	var Amount float64
	err := r.DB.Raw(`SELECT SUM(price) FROM carts WHERE user_id = ?`, userid).Scan(&Amount).Error

	if err != nil {
		return 0.0, nil
	}
	return Amount, nil
}

func (r *Repo) CartFinalPrice(userid uint) (float64, error) {
	var Amount float64
	err := r.DB.Raw(`SELECT SUM(final_price) FROM carts WHERE user_id = ?`, userid).Scan(&Amount).Error

	if err != nil {
		return 0.0, nil
	}
	return Amount, nil
}

func (r *Repo) DisplayCart(userid uint) ([]models.Cart, error) {

	var count int
	if err := r.DB.Raw("SELECT COUNT(*) FROM carts WHERE user_id = ? ", userid).First(&count).Error; err != nil {
		return []models.Cart{}, err
	}

	if count == 0 {
		return []models.Cart{}, nil
	}

	var Cart []models.Cart

	if err := r.DB.Raw("SELECT carts.user_id,users.firstname as user_name,carts.product_id,products.name as product_name,categories.category as category,carts.quantity,carts.price,carts.final_price FROM carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id inner join categories on categories.id = products.category_id where user_id = ?", userid).First(&Cart).Error; err != nil {
		return []models.Cart{}, err
	}

	return Cart, nil
}

func (r *Repo) ViewAddress(id uint) ([]models.AddressRes, error) {
	var Address []models.AddressRes
	query := r.DB.Raw(`SELECT * FROM addresses WHERE user_id = ?`, id).Scan(&Address)
	if query.Error != nil {
		return []models.AddressRes{}, query.Error
	}

	if query.RowsAffected < 1 {
		return []models.AddressRes{}, errors.New("no address found. add new address")
	}

	return Address, nil
}

func (r *Repo) CheckAddressExist(userid uint, address uint) bool {
	var count int
	if err := r.DB.Raw("SELECT COUNT(*) FROM addresses WHERE id = ? AND user_id = ?", address, userid).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (r *Repo) CheckPaymentMethodExist(paymentid uint) bool {
	query := r.DB.Raw(`SELECT * FROM payment_methods WHERE id = ?`, paymentid)
	return query.RowsAffected < 1
}

func (r *Repo) CheckCartExist(userID uint) bool {
	var count int
	if err := r.DB.Raw("SELECT COUNT(*) FROM carts WHERE  user_id = ?", userID).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (r *Repo) OrderFromCart(addressid uint, paymentid, userid uint, price float64) (int, error) {
	var id int
	query := `
    INSERT INTO orders (created_at , user_id , address_id ,payment_method_id,total_price)
    VALUES (NOW(),?, ?, ?,?)
    RETURNING id`
	r.DB.Raw(query, userid, addressid, paymentid, price).Scan(&id)
	return id, nil
}

func (r *Repo) AddOrderProducts(userID uint, orderid int, cart []models.Cart) error {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	query := `
    INSERT INTO order_items (order_id,product_id,user_id,quantity,total_price)
    VALUES (?, ?, ?, ?, ?) `

	for _, v := range cart {
		var productID int
		if err := tx.Raw("SELECT id FROM products WHERE name = $1", v.ProductName).Scan(&productID).Error; err != nil {
			tx.Rollback()
			return errors.New(`something went wrong`)
		}
		if err := tx.Exec(query, orderid, productID, userID, v.Quantity, v.Price).Error; err != nil {
			tx.Rollback()
			return errors.New(`something went wrong`)
		}
	}

	tx.Commit()
	return nil
}

func (r *Repo) UpdateCartAndStockAfterOrder(userID uint, productID int, quantity float64) error {
	err := r.DB.Exec("DELETE FROM carts WHERE user_id = ? and product_id = ?", userID, productID).Error
	if err != nil {
		return errors.New(`something went wrong`)
	}

	err = r.DB.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", quantity, productID).Error
	if err != nil {
		return errors.New(`something went wrong`)
	}

	return nil
}

func (r *Repo) GetUserById(id int) (*models.UserDetailsResponse, error) {
	if r.DB == nil {
		return nil, errors.New("DB is nil")
	}

	var user models.UserDetailsResponse

	result := r.DB.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user)
	if result.Error != nil {
		fmt.Println("Error fetching user:", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no user found with this ID")
	}

	return &user, nil
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

func (r *Repo) CheckOrder(orderid string, userID uint) error {
	var count int
	err := r.DB.Raw("SELECT COUNT(*) FROM order_items WHERE order_id = ? AND user_id = ?", orderid, userID).Scan(&count).Error
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New(`no orders found`)
	}
	return nil
}

func (r *Repo) CancelOrder(orderid, pid string, userID uint) error {
	status := "Cancelled"
	err := r.DB.Exec("UPDATE order_items SET order_status = ?  WHERE order_id = ? AND product_id = ? AND user_id = ?", status, orderid, pid, userID).Error
	if err != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func (r *Repo) UpdateStock(pid int, quantity int) error {
	query := r.DB.Exec(`UPDATE products SET stock = stock + $1 WHERE id = $2`, quantity, pid)
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func (r *Repo) ReturnAmountToWallet(userID uint, orderID, pid string) error {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var amount float64
	query := tx.Raw(`SELECT total_price FROM order_items WHERE product_id = ? AND order_id = ? AND user_id = ?`, pid, orderID, userID).Scan(&amount)
	if query.Error != nil {
		tx.Rollback()
		return errors.New(`something went wrong`)
	}
	query = tx.Exec(`UPDATE users SET wallet = wallet + $1 WHERE id = $2`, amount, userID)
	if query.Error != nil {

		tx.Rollback()
		return errors.New(`something went wrong`)
	}

	if query.RowsAffected == 0 {

		tx.Rollback()
		return errors.New(`no orders found with this id`)
	}

	tx.Commit()

	return nil
}

func (r *Repo) UpdateOrderFinalPrice(orderID int, amount float64) error {
	query := r.DB.Exec(`UPDATE orders SET final_price = final_price - $1 WHERE id = $2`, amount, orderID)
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func (r *Repo) UpdateWallet(wallet float64, userID uint) error {
	query := r.DB.Exec(`UPDATE users SET wallet = ? WHERE id = ?`, wallet, userID)
	if query.Error != nil {
		return errors.New(`something went wrong`)
	}
	return nil
}

func (r *Repo) CancelOrderDetails(userID uint, orderID, pid string) (models.CancelDetails, error) {
	var Details models.CancelDetails
	query := r.DB.Raw(`SELECT order_status,quantity,orders.payment_status,order_items.total_price,order_id FROM order_items INNER JOIN orders ON orders.id =order_items.order_id WHERE order_items.order_id = ? AND order_items.user_id = ? AND order_items.product_id = ?`, orderID, userID, pid).Scan(&Details)
	if query.Error != nil {
		return models.CancelDetails{}, errors.New(`something went wrong`)
	}
	return Details, nil
}

func (r *Repo) GetOrderDetails(userID uint) ([]models.ViewOrderDetails, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var orderDetails []models.OrderDetails
	query := tx.Raw(`
        SELECT orders.id, total_price as final_price, payment_methods.payment_mode AS payment_method, payment_status
        FROM orders
        INNER JOIN payment_methods ON orders.payment_method_id = payment_methods.id
        WHERE user_id = ? ORDER BY orders.id DESC`, userID).Scan(&orderDetails)

	if query.Error != nil {
		tx.Rollback()
		return []models.ViewOrderDetails{}, errors.New(`something went wrong`)
	}

	var fullOrderDetails []models.ViewOrderDetails
	for _, order := range orderDetails {
		var orderProductDetails []models.OrderProductDetails
		query = tx.Raw(`
            SELECT order_items.product_id, products.name AS product_name, order_items.order_status,
                   order_items.quantity, order_items.total_price
            FROM order_items
            INNER JOIN products ON order_items.product_id = products.id
            WHERE order_items.order_id = ? ORDER BY order_id DESC`, order.Id).Scan(&orderProductDetails)

		if query.Error != nil {
			tx.Rollback()
			return []models.ViewOrderDetails{}, errors.New(`something went wrong`)
		}

		fullOrderDetails = append(fullOrderDetails, models.ViewOrderDetails{
			OrderDetails:        order,
			OrderProductDetails: orderProductDetails,
		})
	}

	tx.Commit()
	return fullOrderDetails, nil
}
