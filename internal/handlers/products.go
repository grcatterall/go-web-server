package handlers

import (
	"fmt"
	"net/http"
	"errors"

	"github.com/grcatterall/go-web-server/internal/factories"
	"github.com/grcatterall/go-web-server/internal/repositories"
	"github.com/grcatterall/go-web-server/pkg/utils"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	Repo repositories.ProductRepository
}

func NewProductHandler(repo repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	rf := factories.NewResponseFactory()
	defer rf.ResponseDefer(w)

	products, err := ph.Repo.GetAllProducts()

	if err != nil {
		fmt.Println("Unable to parse products json")
		panic(utils.ErrorResponse{Code: 500, Msg: "Server error"})
	}

	jsonResponse, err := utils.ConvertToJson(products)

	if err != nil {
		fmt.Println("Unable to parse products json")
		panic(utils.ErrorResponse{Code: 500, Msg: "Server error"})
	}

	rf.SuccessResponse(w, 200, jsonResponse)
}

func (ph *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	rf := factories.NewResponseFactory()
	defer rf.ResponseDefer(w)

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		panic(utils.ErrorResponse{Code: 400, Msg: "Missing product id"})
	}

	fmt.Println("Getting product by id")

	var productRepository repositories.ProductRepository = &repositories.ProductRepo{}

	product, err := productRepository.GetProductById(id)

	if err != nil {
		fmt.Println(err)
		if errors.Is(err, repositories.ErrProductNotFound) {
			panic(utils.ErrorResponse{Code: 404, Msg: fmt.Sprintf("Unable to find product with id %s", id)})
		} else {
			panic(utils.ErrorResponse{Code: 500, Msg: "Server error"})
		}
	}

	var response []byte
	var status int

	jsonResponse, err := utils.ConvertToJson(product)

	if err != nil {
		fmt.Println("Unable to parse products json")
		panic(utils.ErrorResponse{Code: 500, Msg: "Server error"})
	}

	response = jsonResponse
	status = 200

	rf.SuccessResponse(w, status, response)
}
