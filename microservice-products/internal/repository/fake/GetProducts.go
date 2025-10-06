package fake

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/models"
	"github.com/google/uuid"
)

// GetProducts is an implement of repository.GetProduct interface
func (f *Fake) GetProducts(_ context.Context, _ string) ([]*models.Product, error) {
	return []*models.Product{
		{
			ID:          uuid.UUID{},
			Name:        "Product 1",
			Description: "Description 1",
			Params:      map[string]any{"param1": "value1"},
			Weight:      13,
			BarCode:     "123213",
		},
		{
			ID:          uuid.UUID{},
			Name:        "Product 2",
			Description: "Description 2",
			Params:      map[string]any{"param2": "value2"},
			Weight:      15,
			BarCode:     "123214",
		},
	}, nil
}
