package grpc

import (
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/service"
	pb "github.com/PIRSON21/grpc-microservices/microservice-warehouses/proto"
)

type WarehouseHandler struct {
	pb.UnimplementedWarehouseServiceServer
	service *service.WarehouseService
}

func NewWarehouseHandler(service *service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{
		service: service,
	}
}
