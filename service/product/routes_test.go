package product

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiofsr/ecom/types"
	"github.com/gorilla/mux"
)

func TestProductServiceHandler(t *testing.T) {
	productStore := &mockProductStore{}
	handler := NewHandler(productStore)

	t.Run("should fail if payload is invalid", func(t *testing.T) {
		payload := types.CreateProductPayload{
			Name:        "Invalid Name",
			Description: "Invalid Description",
			Price:       4.99,
			Quantity:    10,
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleCreateProduct)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should create a new product", func(t *testing.T) {
		payload := types.CreateProductPayload{
			Name:        "Name",
			Description: "Description",
			Price:       4.99,
			Quantity:    10,
			Image:       "image.jpg",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleCreateProduct)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockProductStore struct{}

func (m *mockProductStore) GetProducts() ([]types.Product, error) {
	return nil, nil
}

func (m *mockProductStore) GetProductsByIDs(ids []int) ([]types.Product, error) {
	return nil, nil
}
func (m *mockProductStore) UpdateProduct(types.Product) error {
	return nil
}

func (m *mockProductStore) CreateProduct(product types.Product) error {
	return nil
}
