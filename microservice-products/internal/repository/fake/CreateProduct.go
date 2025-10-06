package fake

import (
	"context"
	"fmt"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/models"
)

// CreateProduct is an implement of repository.PostProduct interface
func (f *Fake) CreateProduct(_ context.Context, product *models.Product) error {
	fmt.Println("NOP CreateProduct", product)
	return nil
}
