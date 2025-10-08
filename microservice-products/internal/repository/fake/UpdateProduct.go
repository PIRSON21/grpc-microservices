package fake

import (
	"context"
	"fmt"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/models"
)

func (r *Fake) UpdateProduct(_ context.Context, prod *models.Product) error {
	fmt.Println("Fake UpdateProduct called")
	fmt.Println(prod)

	return nil
}
