package grpc

import (
	"context"

	pb "github.com/PIRSON21/grpc-microservices/microservice-products/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetProductByID handles gRPC requests to retrieve a product by its ID.
func (h *ProductGRPCHandler) GetProductByID(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	log := h.logger.With().Fields(map[string]any{
		"op": "handler.ProductGRPCHandler.GetProductByID",
		"id": req.Id,
	}).Logger()

	log.Debug().Msg("Received gRPC request to get product by ID")

	if req.Id == "" {
		log.Debug().Msg("Received empty ID in request")
		return nil, status.Error(codes.InvalidArgument, "Product ID cannot be empty")
	}

	products, err := h.service.GetProducts(ctx, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving product")
		return nil, status.Error(codes.Internal, "Error retrieving product")
	}

	if len(products) == 0 {
		log.Debug().Msg("No product found with the given ID")
		return nil, status.Error(codes.NotFound, "Product not found")
	}

	product := products[0]
	resp := &pb.GetProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Weight:      product.Weight,
	}

	return resp, nil
}
