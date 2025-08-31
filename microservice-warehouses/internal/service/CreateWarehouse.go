package service

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/dto"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/dto/db_dto"
	"github.com/sirupsen/logrus"
)

func (s *WarehouseService) CreateWarehouse(ctx context.Context, payload *dto.CreateWarehouseRequest) error {
	log := logrus.WithFields(logrus.Fields{
		"op": "service.WarehouseService.CreateWarehouse",
	})

	log.Debugf("Creating warehouse with payload: %+v", payload)

	warehouse := parsePayloadToDB(payload)

	if err := s.repo.CreateWarehouse(ctx, warehouse); err != nil {
		log.Errorf("Error creating warehouse in repository: %v", err)
		return err
	}

	log.Debug("Warehouse created successfully in repository")

	return nil
}

func parsePayloadToDB(payload *dto.CreateWarehouseRequest) *db_dto.CreateWarehouseDTO {
	dbDTO := db_dto.CreateWarehouseDTO{
		Address: payload.Address,
	}

	return &dbDTO
}
