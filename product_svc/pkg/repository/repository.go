package repository

import (
	"errors"
	"fmt"

	"github.com/14jasimmtp/cityvibe-microservice/product_svc/pkg/domain"
	"github.com/14jasimmtp/cityvibe-microservice/product_svc/pkg/models"
	infRepo "github.com/14jasimmtp/cityvibe-microservice/product_svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) infRepo.Repo {
	return &Repo{DB: db}
}

func (r *Repo) AddProduct(product models.AddProduct) (models.UpdateProduct, error) {
	var dproduct models.UpdateProduct
	var p domain.Product
	result := r.DB.Raw("INSERT INTO products(name,description,category_id,size_id,stock,price,color,image_url) values(?,?,?,?,?,?,?,?)", product.Name, product.Description, product.CategoryID, product.Size, product.Stock, product.Price, product.Color, product.ImageURL).Scan(&p)
	fmt.Println(p)
	if result.Error != nil {
		return models.UpdateProduct{}, result.Error
	}
	query := r.DB.Raw(`SELECT products.id,name,description,categories.category,sizes.size,stock,price,color FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE name = ?`, product.Name).Scan(&dproduct)
	if query.Error != nil {
		return models.UpdateProduct{}, query.Error
	}
	fmt.Println(dproduct)
	return dproduct, nil
}

func (r *Repo) DeleteProduct(id int) error {
	query := r.DB.Exec(`UPDATE products SET deleted = true WHERE id = ?`, id)
	if query.Error != nil {
		return errors.New("no product found to delete")
	}
	return nil
}

func (r *Repo) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	query := r.DB.Raw(`SELECT products.id,name,description,categories.category,sizes.size,stock,color,price,offer_prize FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE deleted = false ORDER BY id ASC`).Scan(&products)
	if query.Error != nil {
		return []models.Product{}, query.Error
	}
	return products, nil
}

func (r *Repo) GetSingleProduct(id int64) (models.Product, error) {
	var product models.Product

	query := r.DB.Raw("SELECT products.id as id,name,description,categories.category,sizes.size,stock,color,price FROM products INNER JOIN categories ON categories.id = products.category_id INNER JOIN sizes ON sizes.id=products.size_id WHERE products.id = ?", id).Scan(&product)
	if product.Name == "" {
		return models.Product{}, errors.New("no products found with this id")
	}

	if query.Error != nil {
		return models.Product{}, errors.New("something went wrong")
	}

	return product, nil
}
