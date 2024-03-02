package infRepo

import "github.com/14jasimmtp/cityvibe-microservice/product_svc/pkg/models"



type Repo interface {
	AddProduct(product models.AddProduct) (models.UpdateProduct, error)
	DeleteProduct(id int) error
	GetAllProducts() ([]models.Product, error)
	GetSingleProduct(id int64) (models.Product, error)
}
