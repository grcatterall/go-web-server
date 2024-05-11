package main

import (
	"fmt"
	"net/http"

	"example.com/web-server/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	port := ":80"
	domain := "http://localhost"

	r.HandleFunc("/products/", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/product/{id}", handlers.GetProductById).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "\n")

		fmt.Fprint(w)
	})

	fmt.Printf("Go running at %s", domain+port)
	http.ListenAndServe(port, r)
}
