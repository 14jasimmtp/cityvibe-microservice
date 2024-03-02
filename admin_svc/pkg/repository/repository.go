package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/14jasimmtp/cityvibe-microservices/admin-svc/pkg/models"
	interfaceRepo "github.com/14jasimmtp/cityvibe-microservices/admin-svc/pkg/repository/interface"
	"github.com/14jasimmtp/cityvibe-microservices/admin-svc/utils"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) interfaceRepo.Repo {
	return &Repo{DB: db}
}

func (r *Repo) AdminLogin(adminDetails models.AdminLogin) (models.Admin, error) {
	var details models.Admin
	if err := r.DB.Raw("SELECT * FROM admins WHERE email=?", adminDetails.Email).Scan(&details).Error; err != nil {
		return models.Admin{}, err
	}
	return details, nil
}

func (r *Repo) GetAllUsers() ([]models.UserDetailsResponse, error) {
	var users []models.UserDetailsResponse
	result := r.DB.Raw("SELECT id,email,firstname,lastname,phone,blocked,wallet FROM users").Scan(&users)
	if result.Error != nil {
		fmt.Println("data fetching error")
		return []models.UserDetailsResponse{}, result.Error
	}

	return users, nil
}

func (r *Repo) BlockUserByID(user models.UserDetailsResponse) error {
	result := r.DB.Exec("UPDATE users SET blocked = ? WHERE id = ?", user.Blocked, user.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repo) UnBlockUserByID(user models.UserDetailsResponse) error {
	result := r.DB.Exec("UPDATE users SET blocked = ? WHERE id = ?", user.Blocked, user.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repo) GetAllOrderDetailsBrief() ([]models.ViewAdminOrderDetails, error) {

	var orderDatails []models.AdminOrderDetails
	query := r.DB.Raw("SELECT orders.user_id,orders.id, total_price as final_price, payment_methods.payment_mode AS payment_method, payment_status FROM orders INNER JOIN payment_methods ON orders.payment_method_id=payment_methods.id  ORDER BY orders.id DESC").Scan(&orderDatails)
	if query.Error != nil {
		return []models.ViewAdminOrderDetails{}, errors.New(`something went wrong`)
	}
	var fullOrderDetails []models.ViewAdminOrderDetails
	for _, ok := range orderDatails {
		var OrderProductDetails []models.OrderProductDetails
		r.DB.Raw("SELECT order_items.product_id,products.name AS product_name,order_items.order_status,order_items.quantity,order_items.total_price FROM order_items INNER JOIN products ON order_items.product_id = products.id WHERE order_items.order_id = ? ORDER BY order_id DESC", ok.Id).Scan(&OrderProductDetails)
		fullOrderDetails = append(fullOrderDetails, models.ViewAdminOrderDetails{OrderDetails: ok, OrderProductDetails: OrderProductDetails})
	}
	return fullOrderDetails, nil

}

func (r *Repo) GetSingleOrderDetails(orderID string) ([]models.OrderProductDetails, error) {
	var Order []models.OrderProductDetails
	query := r.DB.Raw(`SELECT product_id,products.name AS product_name,order_status,quantity,Total_price FROM order_items INNER JOIN products ON product_id=products.id WHERE order_id = ?`, orderID).Scan(&Order)
	if query.Error != nil {
		return []models.OrderProductDetails{}, query.Error
	}
	return Order, nil
}

func (r *Repo) DashBoardUserDetails() (models.DashBoardUser, error) {
	var userDetails models.DashBoardUser
	err := r.DB.Raw("SELECT COUNT(*) FROM users").Scan(&userDetails.TotalUsers).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	err = r.DB.Raw("SELECT id FROM users WHERE blocked=true").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	return userDetails, nil
}

func (r *Repo) DashBoardProductDetails() (models.DashBoardProduct, error) {
	var productDetails models.DashBoardProduct
	err := r.DB.Raw("SELECT COUNT(*) FROM products").Scan(&productDetails.TotalProducts).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}
	err = r.DB.Raw("SELECT id FROM products WHERE stock=0").Scan(&productDetails.OutofStockProductID).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}
	err = r.DB.Raw("SELECT id FROM products WHERE stock<=5").Scan(&productDetails.LowStockProductsID).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}
	return productDetails, nil
}

func (r *Repo) TotalRevenue() (models.DashboardRevenue, error) {
	var revenueDetails models.DashboardRevenue
	startTime := time.Now().AddDate(0, 0, -1)
	endTime := time.Now()
	err := r.DB.Raw("SELECT COALESCE(SUM(total_price),0) FROM orders WHERE payment_status = 'paid' AND created_at >=? AND created_at <=?", startTime, endTime).Scan(&revenueDetails.TodayRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	startTime, endTime = utils.CalcualtePeriodDate("monthly")
	err = r.DB.Raw("SELECT COALESCE (SUM(total_price),0) FROM orders WHERE payment_status = 'paid' AND created_at >=? AND created_at <=?", startTime, endTime).Scan(&revenueDetails.MonthRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	startTime, endTime = utils.CalcualtePeriodDate("yearly")
	err = r.DB.Raw("SELECT COALESCE (SUM(total_price),0) FROM orders WHERE payment_status = 'paid' AND created_at >=? AND created_at <=?", startTime, endTime).Scan(&revenueDetails.YearRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	return revenueDetails, nil
}

func (r *Repo) AmountDetails() (models.DashboardAmount, error) {
	var amountDetails models.DashboardAmount
	err := r.DB.Raw("SELECT COALESCE (SUM(total_price),0) FROM orders WHERE payment_status = 'paid' ").Scan(&amountDetails.CreditedAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}
	err = r.DB.Raw("SELECT COALESCE(SUM(total_price),0) FROM orders WHERE payment_status = 'not paid' ").Scan(&amountDetails.PendingAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}
	return amountDetails, nil
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

func (r *Repo) DashBoardOrder() (models.DashboardOrder, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var orderDetail models.DashboardOrder
	err := tx.Raw("SELECT COUNT(*) FROM order_items WHERE order_status= 'Delivered'").Scan(&orderDetail.DeliveredOrderProducts).Error
	if err != nil {
		tx.Rollback()
		return models.DashboardOrder{}, err
	}

	err = tx.Raw("SELECT COUNT(*) FROM order_items WHERE order_status='pending' OR order_status = 'processing'").Scan(&orderDetail.PendingOrderProducts).Error
	if err != nil {
		tx.Rollback()
		return models.DashboardOrder{}, err
	}

	err = tx.Raw("SELECT COUNT(*) FROM order_items WHERE order_status = 'Cancelled' OR order_status = 'returned'").Scan(&orderDetail.CancelledOrderProducts).Error
	if err != nil {
		tx.Rollback()
		return models.DashboardOrder{}, err
	}

	err = tx.Raw("SELECT COUNT(*) FROM order_items").Scan(&orderDetail.TotalOrderItems).Error
	if err != nil {
		tx.Rollback()
		return models.DashboardOrder{}, err
	}

	err = tx.Raw("SELECT COALESCE(SUM(quantity), 0) FROM order_items").Scan(&orderDetail.TotalOrderQuantity).Error
	if err != nil {
		tx.Rollback()
		return models.DashboardOrder{}, err
	}

	tx.Commit()
	return orderDetail, nil
}