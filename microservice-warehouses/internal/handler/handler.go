package handler

import (
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/service"
)

// WarehouseHandler handles HTTP requests related to warehouses.
type WarehouseHandler struct {
	service *service.WarehouseService
}

// NewWarehouseHandler creates a new instance of WarehouseHandler with the provided WarehouseService.
func NewWarehouseHandler(s *service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{service: s}
}
