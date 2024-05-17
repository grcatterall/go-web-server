package repositories

import (
	"fmt"
	"errors"

	"github.com/grcatterall/go-web-server/internal/models"
)

var products = []models.Product{
	{ID: "172nw-4719enqlos-4710", Name: "Sun Glasses", Price: 12.99, Description: "Nice shades"},
	{ID: "172nw-4719enqlos-4711", Name: "Jumper", Price: 30.00, Description: "Lovely Jumper"},
	{ID: "172nw-4719enqlos-4712", Name: "Sun Dial", Price: 1099.50, Description: "Tell the time"},
}

var ErrProductNotFound = errors.New("product not found")

type ProductRepo struct{}

func (r *ProductRepo) GetAllProducts() ([]models.Product, error) {
	fmt.Println("Getting all products from within repository")

	return products, nil
}

func (r *ProductRepo) GetProductById(id string) (models.Product, error) {
	fmt.Println("Getting product by id from within repository")

	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}

	return models.Product{}, ErrProductNotFound
}

func (r *ProductRepo) GetProductByName(name string) (models.Product, error) {
	fmt.Println("Getting product by name from within repository")

	return products[1], nil
}

func (r *ProductRepo) CreateProduct(product models.Product) (models.Product, error) {
	fmt.Println("Creating product from within repository")

	return products[1], nil
}

func (r *ProductRepo) UpdateProduct(id string, product models.Product) (models.Product, error) {
	fmt.Println("Updating from within repository")

	return products[0], nil
}

func (r *ProductRepo) DeleteProduct(id string) (bool, error) {
	fmt.Println("Deleting product from within repository")

	return true, nil
}

