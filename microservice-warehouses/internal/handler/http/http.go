package http

import (
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/service"
)

// WarehouseHTTPHandler handles HTTP requests related to warehouses.
type WarehouseHTTPHandler struct {
	service *service.WarehouseService
}

// NewWarehouseHandler creates a new instance of WarehouseHTTPHandler with the provided WarehouseService.
func NewWarehouseHandler(s *service.WarehouseService) *WarehouseHTTPHandler {
	return &WarehouseHTTPHandler{service: s}
}
