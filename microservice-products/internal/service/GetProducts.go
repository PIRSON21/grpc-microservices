package service

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/dto"
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/models"
	"github.com/sirupsen/logrus"
)

// GetProducts retrieves a list of products based on the provided ID.
func (s *ProductService) GetProducts(ctx context.Context, id string) ([]*dto.GetProductResponse, error) {
	log := logrus.WithField("method", "GetProducts")
	log.Infof("Retrieving products with ID: %s", id)

	products, err := s.repo.GetProducts(ctx, id)
	if err != nil {
		log.Errorf("Failed to retrieve products: %v", err)
		return nil, err
	}

	var responses []*dto.GetProductResponse
	for _, product := range products {
		response := parseProductToGetResponse(product)
		responses = append(responses, response)
	}

	log.Infof("Retrieved %d products", len(responses))

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
