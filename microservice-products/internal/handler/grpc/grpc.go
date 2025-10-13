package grpc

import (
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/service"
	pb "github.com/PIRSON21/grpc-microservices/microservice-products/proto"
	"github.com/rs/zerolog"
)

// ProductGRPCHandler handles gRPC requests for products.
type ProductGRPCHandler struct {
	logger *zerolog.Logger
	pb.UnimplementedProductServiceServer
	service *service.ProductService
}

// NewProductGRPCHandler creates a new ProductGRPCHandler.
func NewProductGRPCHandler(service *service.ProductService, logger *zerolog.Logger) *ProductGRPCHandler {
	return &ProductGRPCHandler{
		service: service,
		logger:  logger,
	}
}
