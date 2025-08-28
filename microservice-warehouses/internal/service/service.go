package service

import "github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/repository"

// WarehouseService provides methods to manage warehouses.
type WarehouseService struct {
	repo repository.WarehouseRepository
}

// NewWarehouseService creates a new instance of WarehouseService.
func NewWarehouseService(repo repository.WarehouseRepository) *WarehouseService {
	return &WarehouseService{
		repo: repo,
	}
}
