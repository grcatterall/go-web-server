package handlers

import (
    // "bytes"
    // "errors"
    "net/http"
    "net/http/httptest"
    "testing"

    // "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/mock"
    "github.com/grcatterall/go-web-server/internal/models"
    // "github.com/grcatterall/go-web-server/internal/repositories"
    // "github.com/grcatterall/go-web-server/pkg/utils"
    // "github.com/grcatterall/go-web-server/internal/factories"
    "github.com/grcatterall/go-web-server/mocks"
)

var products = []models.Product{
    {ID: "172nw-4719enqlos-4710", Name: "Sun Glasses", Price: 12.99, Description: "Nice shades"},
    {ID: "172nw-4719enqlos-4711", Name: "Jumper", Price: 30.00, Description: "Lovely Jumper"},
    {ID: "172nw-4719enqlos-4712", Name: "Sun Dial", Price: 1099.50, Description: "Tell the time"},
}

func TestGetProducts(t *testing.T) {
	mockRepository := new(mocks.ProductRepository)
	mockRepository.On("GetAllProducts").Return(products, nil)

    handler := NewProductHandler(mockRepository)

	req, err := http.NewRequest("GET", "/products", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	httpHandler := http.HandlerFunc(handler.GetProducts)

	// Act
	httpHandler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Sun Glasses")
	mockRepository.AssertExpectations(t)
}