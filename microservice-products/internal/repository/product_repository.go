package repository

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/models"
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/repository/fake"
)

// GetProduct defines the interface for retrieving products.
type GetProduct interface {
	// GetProducts retrieves a list of products based on the provided ID.
	GetProducts(context.Context, string) ([]*models.Product, error)
}

// PostProduct defines the interface for creating and deleting products.
type PostProduct interface {
	// CreateProduct adds a new product to the repository.
	CreateProduct(context.Context, *models.Product) error
	// DeleteProduct removes a product from the repository by its ID.
	DeleteProduct(string) error
	// UpdateProduct modifies an existing product in the repository.
	UpdateProduct(context.Context, *models.Product) error
}

// ProductRepository defines the interface for product data operations.
type ProductRepository interface {
	GetProduct
	PostProduct
	// Close releases any resources held by the repository.
	Close() error
}

// NewProductRepository creates a new instance of a ProductRepository.
func NewProductRepository() ProductRepository {
	return fake.NewFakeProductRepository()
}
