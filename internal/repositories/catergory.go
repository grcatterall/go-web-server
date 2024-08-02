package repositories

import (
	"errors"
	"fmt"
	"log"

	"github.com/grcatterall/go-web-server/internal/models"
	"github.com/grcatterall/go-web-server/pkg/utils"
)


var ErrCategoryNotFound = errors.New("product not found")

type CategoryRepo struct{}

func (c CategoryRepo) GetAllCategories() ([]models.Category, error) {
	conn := utils.DbConnection()

	rows, err := conn.Query("SELECT * FROM categories")

	var categories []models.Category

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category models.Category

		if err := rows.Scan(&category.ID, &category.Name, &category.Description, &category.Slug); err != nil {
			log.Fatal(err)
		}

		categories = append(categories, category)
		fmt.Println(category.Name)
	}

	rows.Close()

	conn.Close()

	return categories, nil
}