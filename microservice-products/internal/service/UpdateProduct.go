package service

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/dto"
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/models"
)

func (s *ProductService) UpdateProduct(ctx context.Context, request *dto.UpdateProductRequest) error {
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
		return err
	}

	// TODO: доделать метод UpdateProduct (в service и handler)

	log.Info().Msgf("Product updated with ID: %s", product.ID)

	return nil
}

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
