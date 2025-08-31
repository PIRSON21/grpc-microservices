package grpc

import (
	"context"

	pb "github.com/PIRSON21/grpc-microservices/microservice-warehouses/proto"
	"github.com/sirupsen/logrus"
)

// GetWarehouseByID handles gRPC requests to retrieve a warehouse by its ID.
func (h *WarehouseHandler) GetWarehouseByID(ctx context.Context, req *pb.GetWarehouseRequest) (*pb.WarehouseResponse, error) {
	log := logrus.WithFields(logrus.Fields{
		"op": "handler.grpc.WarehouseHandler.GetWarehouseByID",
		"id": req.Id,
	})

	log.Debug("Received gRPC request to get warehouse by ID")

	if req.Id == "" {
		log.Warn("Received empty ID in request")
		return &pb.WarehouseResponse{}, nil
	}

	warehouses, err := h.service.GetWarehouses(ctx, req.Id)
	if err != nil {
		log.Errorf("Error retrieving warehouse: %v", err)
		return nil, err
	}

	if len(warehouses) == 0 {
		log.Warn("No warehouse found with the given ID")
		return &pb.WarehouseResponse{}, nil
	}

	warehouse := warehouses[0]
	resp := &pb.WarehouseResponse{
		Id:      warehouse.ID,
		Address: warehouse.Address,
	}

	log.Debug("Successfully retrieved warehouse by ID")

	return resp, nil
}
