package service

import (
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/repository"
	"github.com/rs/zerolog"
)

// ProductService struct provides methods to interact with product data.
type ProductService struct {
	repo   repository.ProductRepository
	logger *zerolog.Logger
}

// NewProductService creates a new instance of ProductService.
func NewProductService(repo repository.ProductRepository, logger *zerolog.Logger) *ProductService {
	return &ProductService{repo: repo, logger: logger}
}
