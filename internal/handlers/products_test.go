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

func TestGetProducts(t *testing.T) {
	mockRepository := new(mocks.MockProductRepository)
	mockRepository.On("GetAllProducts").Return([]models.Product{
		{ID: "172nw-4719enqlos-4710", Name: "Sun Glasses", Price: 12.99, Description: "Nice shades"},
		{ID: "172nw-4719enqlos-4711", Name: "Jumper", Price: 30.00, Description: "Lovely Jumper"},
		{ID: "172nw-4719enqlos-4712", Name: "Sun Dial", Price: 1099.50, Description: "Tell the time"},
	}, nil)

	req, err := http.NewRequest("GET", "/products", nil)
    assert.NoError(t, err)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(GetProducts)

    // Act
    handler.ServeHTTP(rr, req)

    // Assert
    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Contains(t, rr.Body.String(), "Sun Glasses")
}