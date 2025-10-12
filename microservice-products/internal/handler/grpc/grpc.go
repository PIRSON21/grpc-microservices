package grpc

import (
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/service"
	pb "github.com/PIRSON21/grpc-microservices/microservice-products/proto"
)

// ProductGRPCHandler handles gRPC requests for products.
type ProductGRPCHandler struct {
	pb.UnimplementedProductServiceServer
	service *service.ProductService
}

// NewProductGRPCHandler creates a new ProductGRPCHandler.
func NewProductGRPCHandler(service *service.ProductService) *ProductGRPCHandler {
	return &ProductGRPCHandler{
		service: service,
	}
}
