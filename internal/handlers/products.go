package handlers

import (
	"fmt"
	"net/http"

	"github.com/grcatterall/go-web-server/internal/factories"
	"github.com/grcatterall/go-web-server/internal/models"
	"github.com/grcatterall/go-web-server/pkg/utils"

	"github.com/gorilla/mux"
)

var products = []models.Product{
	{ID: "172nw-4719enqlos-4710", Name: "Sun Glasses", Price: 12.99, Description: "Nice shades"},
	{ID: "172nw-4719enqlos-4711", Name: "Jumper", Price: 30.00, Description: "Lovely Jumper"},
	{ID: "172nw-4719enqlos-4712", Name: "Sun Dial", Price: 1099.50, Description: "Tell the time"},
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	rf := factories.NewResponseFactory()
	defer rf.ResponseDefer(w)
	jsonResponse, err := utils.ConvertToJson(products)

	if err != nil {
		fmt.Println("Unable to parse products json")
		panic(utils.ErrorResponse{Code: 500, Msg: "Server error"})
	}

	rf.SuccessResponse(w, 200, jsonResponse)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	rf := factories.NewResponseFactory()
	defer rf.ResponseDefer(w)

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		panic(utils.ErrorResponse{Code: 400, Msg: "Missing product id"})
	}

	fmt.Println("Getting product by id")

	var reqProduct models.Product

	for i, product := range products {
		if product.ID == id {
			reqProduct = products[i]
		}
	}

	var response []byte
	var status int

	if reqProduct.GetName() != "" {
		jsonResponse, err := utils.ConvertToJson(reqProduct)

		if err != nil {
			fmt.Println("Unable to parse products json")
			panic(utils.ErrorResponse{Code: 500, Msg: "Server error"})
		}

		response = jsonResponse
		status = 200
	} else {
		panic(utils.ErrorResponse{Code: 404, Msg: fmt.Sprintf("Unable to find product with id %s", id)})
	}

	rf.SuccessResponse(w, status, response)
}
