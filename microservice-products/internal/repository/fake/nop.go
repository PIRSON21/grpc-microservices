package fake

import "github.com/PIRSON21/grpc-microservices/microservice-products/internal/config"

type Fake struct{}

func NewFakeProductRepository(_ *config.DBConfig) *Fake {
	return &Fake{}
}

// Close implements the repository.ProductRepository interface.
func (f *Fake) Close() error {
	return nil
}
