package handlers

import (
	"net/http"

	"example.com/web-server/internal/models"
	"example.com/web-server/pkg/utils"


	"github.com/gorilla/mux"
)

var products = []models.Product{
	{ID: "172nw-4719enqlos-4710", Name: "Sun Glasses", Price: 12.99, Description: "Nice shades"},
	{ID: "172nw-4719enqlos-4711", Name: "Jumper", Price: 30.00, Description: "Lovely Jumper"},
	{ID: "172nw-4719enqlos-4712", Name: "Sun Dial", Price: 1099.50, Description: "Tell the time"},
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	jsonResponse, err := utils.ConvertToJson(products)

	if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }

	utils.JsonResponse(w, jsonResponse)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var reqProduct models.Product

	for i, product := range products {
		if product.ID == id {
			reqProduct = products[i]
		}
	}

	jsonResponse, err := utils.ConvertToJson(reqProduct)

	if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }

	utils.JsonResponse(w, jsonResponse)
}