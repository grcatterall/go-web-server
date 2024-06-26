package main

import (
	"fmt"
	"net/http"

	"github.com/grcatterall/go-web-server/internal/handlers"
	"github.com/grcatterall/go-web-server/internal/repositories"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	port := ":80"
	domain := "http://localhost"

	productHandler := initHandlers()

	r.HandleFunc("/products/", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/product/{id}", productHandler.GetProductById).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "\n")

		fmt.Fprint(w)
	})

	fmt.Printf("Go running at %s", domain+port)
	fmt.Println()
	http.ListenAndServe(port, r)
}

func initHandlers() handlers.ProductHandler {
	var productRepository repositories.ProductRepository = &repositories.ProductRepo{}

	productHandler := handlers.NewProductHandler(productRepository)

	return *productHandler
}