package service

import (
	"context"
	"errors"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/dto"
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/models"
)

var ValidationError = errors.New("validation error")

// CreateProduct creates a new product based on the provided request DTO.
func (s *ProductService) CreateProduct(ctx context.Context, request *dto.CreateProductRequest) (*dto.CreateProductResponse, error) {
	log := s.logger.With().Fields(map[string]any{
		"method": "CreateProduct",
	}).Logger()
	log.Info().Msgf("Creating product with name: %s", request.Name)

	product := parseRequestToProduct(request)

	// Create the product in the repository
	err := s.repo.CreateProduct(ctx, product)
	if err != nil {
		log.Error().Msgf("Failed to create product: %v", err)
		return nil, err
	}

	log.Info().Msgf("Product created with ID: %s", product.ID)

	response := parseProductToResponse(product)

	return response, nil
}

// parseRequestToProduct converts a dto.CreateProductRequest to a models.Product.
func parseRequestToProduct(request *dto.CreateProductRequest) *models.Product {
	return &models.Product{
		Name:        request.Name,
		Description: request.Description,
		Params:      copyMap(request.Params),
		Weight:      request.Weight,
		BarCode:     request.BarCode,
	}
}

// copyMap creates a shallow copy of a map[string]any.
func copyMap(original map[string]any) map[string]any {
	if original == nil {
		return nil
	}
	copied := make(map[string]any, len(original))
	for k, v := range original {
		copied[k] = v
	}
	return copied
}

// parseProductToResponse converts a Product model to a CreateProductResponse DTO.
func parseProductToResponse(product *models.Product) *dto.CreateProductResponse {
	return &dto.CreateProductResponse{
		ID:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		Params:      copyMap(product.Params),
		Weight:      product.Weight,
		BarCode:     product.BarCode,
	}
}
