package grpc

import (
	"context"

	pb "github.com/PIRSON21/grpc-microservices/microservice-warehouses/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetWarehouseByID handles gRPC requests to retrieve a warehouse by its ID.
func (h *WarehouseHandler) GetWarehouseByID(ctx context.Context, req *pb.GetWarehouseRequest) (*pb.WarehouseResponse, error) {
	log := logrus.WithFields(logrus.Fields{
		"op": "handler.grpc.WarehouseHandler.GetWarehouseByID",
		"id": req.Id,
	})

	log.Debug("Received gRPC request to get warehouse by ID")

	if req.Id == "" {
		log.Debug("Received empty ID in request")
		return nil, status.Error(codes.InvalidArgument, "ID cannot be empty")
	}

	warehouses, err := h.service.GetWarehouses(ctx, req.Id)
	if err != nil {
		log.Errorf("Error retrieving warehouse: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to retrieve warehouse: %v", err)
	}

	if len(warehouses) == 0 {
		log.Debug("No warehouse found with the given ID")
		return nil, status.Error(codes.NotFound, "Warehouse not found")
	}

	warehouse := warehouses[0]
	resp := &pb.WarehouseResponse{
		Id:      warehouse.ID,
		Address: warehouse.Address,
	}

	log.Debug("Successfully retrieved warehouse by ID")

	return resp, nil
}
