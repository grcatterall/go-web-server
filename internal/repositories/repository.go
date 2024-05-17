package repositories

import "github.com/grcatterall/go-web-server/internal/models"

type ProductRepository interface {
    GetAllProducts() ([]models.Product, error)
	GetProductById(id string) (models.Product, error)
	GetProductByName(name string) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(id string, product models.Product) (models.Product, error)
	DeleteProduct(id string) (bool, error)
}