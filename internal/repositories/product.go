package repositories

import (
	"fmt"
	"log"
	"errors"
	"database/sql"

	"github.com/grcatterall/go-web-server/internal/models"
	_ "github.com/lib/pq"
)

// var products = []models.Product{
// 	{ID: "172nw-4719enqlos-4710", Name: "Sun Glasses", Price: 12.99, Description: "Nice shades"},
// 	{ID: "172nw-4719enqlos-4711", Name: "Jumper", Price: 30.00, Description: "Lovely Jumper"},
// 	{ID: "172nw-4719enqlos-4712", Name: "Sun Dial", Price: 1099.50, Description: "Tell the time"},
// }

var ErrProductNotFound = errors.New("product not found")

type ProductRepo struct{}

func dbConnection() *sql.DB {

	connectionStr := "postgres://postgres:password@db:5432/go-ecom?sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	return conn
}

func (r *ProductRepo) GetAllProducts() ([]models.Product, error) {
	fmt.Println("Getting all products from within repository")

	var products = []models.Product{}

	conn := dbConnection()

	rows, err := conn.Query("SELECT * FROM products")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product

		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			log.Fatal(err)
		}

		products = append(products, product)

		fmt.Println(product.GetName())
	}

	rows.Close()

	conn.Close()


	return products, nil
}

func (r *ProductRepo) GetProductById(id string) (models.Product, error) {
	fmt.Println("Getting product by id from within repository")

	var products = []models.Product{}

	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}

	return models.Product{}, ErrProductNotFound
}

func (r *ProductRepo) GetProductByName(name string) (models.Product, error) {
	fmt.Println("Getting product by name from within repository")
	
	var products = []models.Product{}

	return products[1], nil
}

func (r *ProductRepo) CreateProduct(product models.Product) (models.Product, error) {
	fmt.Println("Creating product from within repository")
	
	var products = []models.Product{}

	return products[1], nil
}

func (r *ProductRepo) UpdateProduct(id string, product models.Product) (models.Product, error) {
	fmt.Println("Updating from within repository")
	
	var products = []models.Product{}

	return products[0], nil
}

func (r *ProductRepo) DeleteProduct(id string) (bool, error) {
	fmt.Println("Deleting product from within repository")

	return true, nil
}

