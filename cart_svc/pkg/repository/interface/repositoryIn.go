package interfaceRepo

import "github.com/14jasimmtp/cityvibe-microservice/cart_svc/pkg/models"


type Repo interface {
	AddToCart(pid int, userid uint, productAmount float64) error
	DisplayCart(userid uint) ([]models.Cart, error)
	RemoveProductFromCart(pid int, userid uint) error
	CheckProductExistInCart(userId uint, pid string) (bool, error)
	CheckSingleProduct(id string) (models.Product, error)
	GetCartProductAmountFromID(pid string) (float64, error)
	TotalPrizeOfProductInCart(userID uint, productID string) (float64, error)
	UpdateCart(quantity int, price float64, userID uint, product_id string) error
	CartTotalAmount(userid uint) (float64, error)
	CheckCartStock(pid int) error
}
