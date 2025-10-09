package service

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/dto"
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/models"
)

// UpdateProduct updates an existing product based on the provided dto.UpdateProductRequest.
func (s *ProductService) UpdateProduct(ctx context.Context, request *dto.UpdateProductRequest) (*dto.CreateProductResponse, error) {
	log := s.logger.With().Fields(map[string]any{
		"method": "UpdateProduct",
	}).Logger()
	log.Info().Msgf("Updating product with ID: %s", request.ID)

	// Parse DTO to model
	product := parseUpdateRequestToProduct(request)

	// Update the product in the repository
	err := s.repo.UpdateProduct(ctx, product)
	if err != nil {
		log.Error().Msgf("Failed to update product: %v", err)
		return nil, err
	}

	// Retrieve the updated product to return
	updatedProduct, err := s.repo.GetProducts(ctx, product.ID.String())
	if err != nil {
		log.Error().Msgf("Failed to retrieve updated product: %v", err)
		return nil, err
	}

	response := parseProductToResponse(updatedProduct[0])

	log.Info().Msgf("Product updated with ID: %s", product.ID)

	return response, nil
}

// parseUpdateRequestToProduct converts a dto.UpdateProductRequest to a models.Product.
func parseUpdateRequestToProduct(request *dto.UpdateProductRequest) *models.Product {
	return &models.Product{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		Params:      copyMap(request.Params),
		Weight:      request.Weight,
		BarCode:     request.BarCode,
	}
}
