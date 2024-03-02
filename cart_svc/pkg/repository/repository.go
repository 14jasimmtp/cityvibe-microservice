package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/14jasimmtp/cityvibe-microservice/cart_svc/pkg/models"
	interfaceRepo "github.com/14jasimmtp/cityvibe-microservice/cart_svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(r *gorm.DB) interfaceRepo.Repo {
	return &Repo{DB: r}
}

func (r *Repo) AddToCart(pid int, userid uint, productAmount float64) error {
	query := r.DB.Exec(`INSERT INTO carts (user_id,product_id,quantity,price) VALUES (?,?,?,?)`, userid, pid, 1, productAmount)
	if query.Error != nil {
		return query.Error
	}

	return nil
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

func (r *Repo) RemoveProductFromCart(pid int, userid uint) error {
	query := r.DB.Exec(`DELETE FROM carts WHERE product_id = ? AND user_id = ?`, pid, userid)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New(`no products found in cart`)
	}

	return nil
}

func (r *Repo) CheckProductExistInCart(userId uint, pid string) (bool, error) {
	var count int
	query := r.DB.Raw(`SELECT COUNT(*) FROM carts WHERE user_id = ? AND product_id = ?`, userId, pid).Scan(&count)
	if query.Error != nil {
		return false, errors.New(`something went wrong`)
	}
	fmt.Println(count)

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *Repo) CheckSingleProduct(id string) (models.Product, error) {
	var product models.Product
	idint, err := strconv.Atoi(id)
	if err != nil {
		return models.Product{}, errors.New("error while converting id to int")
	}

	query := r.DB.Raw("SELECT products.id as id,name,description,categories.category,sizes.size,stock,color,price FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE products.id = ?", idint).Scan(&product)
	if product.Name == "" {
		return models.Product{}, errors.New("no products found with this id")
	}

	if query.Error != nil {
		return models.Product{}, errors.New("something went wrong")
	}

	return product, nil
}

func (r *Repo) GetCartProductAmountFromID(pid string) (float64, error) {
	var price struct {
		Price      float64
		OfferPrize float64
	}

	if err := r.DB.Raw("select price,offer_prize from products where id = ?", pid).Scan(&price).Error; err != nil {
		return 0.0, err
	}
	if price.OfferPrize != 0 {
		return price.OfferPrize, nil
	}
	return price.Price, nil
}

func (r *Repo) TotalPrizeOfProductInCart(userID uint, productID string) (float64, error) {

	var totalPrice float64
	if err := r.DB.Raw("select sum(price) as total_price from carts where user_id = ? and product_id = ?", userID, productID).Scan(&totalPrice).Error; err != nil {
		return 0.0, err
	}
	return totalPrice, nil
}

func (r *Repo) UpdateCart(quantity int, price float64, userID uint, product_id string) error {

	if err := r.DB.Exec("update carts set quantity = quantity + $1, price = $2 where user_id = $3 and product_id = $4", quantity, price, userID, product_id).Error; err != nil {
		return err
	}

	return nil

}

func (r *Repo) CheckCartStock(pid int) error {
	var stock int
	r.DB.Raw(`SELECT stock from products WHERE id = ?`, pid).Scan(&stock)
	if stock < 1 {
		return errors.New("product out of stock")
	}
	return nil
}

func (r *Repo) CartTotalAmount(userid uint) (float64, error) {
	var Amount float64
	err := r.DB.Raw(`SELECT SUM(price) FROM carts WHERE user_id = ?`, userid).Scan(&Amount).Error

	if err != nil {
		return 0.0, nil
	}
	return Amount, nil
}
