package service

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/dto"
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/models"
)

// GetProducts retrieves a list of products based on the provided ID.
func (s *ProductService) GetProducts(ctx context.Context, id string) ([]*dto.GetProductResponse, error) {
	log := s.logger.With().Fields(map[string]any{
		"method": "GetProducts",
	}).Logger()
	log.Info().Msgf("Retrieving products with ID: %s", id)

	products, err := s.repo.GetProducts(ctx, id)
	if err != nil {
		log.Error().Msgf("Failed to retrieve products: %v", err)
		return nil, err
	}

	var responses []*dto.GetProductResponse
	for _, product := range products {
		response := parseProductToGetResponse(product)
		responses = append(responses, response)
	}

	log.Info().Msgf("Retrieved %d products", len(responses))

	return responses, nil
}

// parseProductToGetResponse converts a Product model to a GetProductResponse DTO.
func parseProductToGetResponse(product *models.Product) *dto.GetProductResponse {
	return &dto.GetProductResponse{
		ID:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		Params:      copyMap(product.Params),
		Weight:      product.Weight,
		BarCode:     product.BarCode,
	}
}
