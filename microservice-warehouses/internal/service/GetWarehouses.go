package service

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/dto"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/dto/db_dto"
	"github.com/sirupsen/logrus"
)

// GetWarehouses retrieves warehouse information based on the provided ID.
func (s *WarehouseService) GetWarehouses(ctx context.Context, id string) ([]*dto.ResponseWarehousesDTO, error) {
	log := logrus.WithFields(logrus.Fields{
		"op": "service.WarehouseService.GetWarehouses",
		"id": id,
	})

	log.Debug("Fetching warehouses from the database")

	warehouses, err := s.repo.GetWarehouses(ctx, id)
	if err != nil {
		log.Errorf("Error fetching warehouses: %v", err)
		return nil, err
	}

	res := parseWarehousesToResponseDTO(warehouses)

	log.WithField("count", len(res)).Debug("Fetched and parsed warehouses successfully")

	return res, nil
}

func parseWarehousesToResponseDTO(warehouses []*db_dto.WarehouseDTO) []*dto.ResponseWarehousesDTO {
	log := logrus.WithFields(logrus.Fields{
		"op": "service.parseWarehousesToResponseDTO",
	})

	log.Debug("Parsing warehouses to ResponseWarehousesDTO")

	var responseWarehouses []*dto.ResponseWarehousesDTO

	for _, wh := range warehouses {
		responseWarehouse := &dto.ResponseWarehousesDTO{
			ID:      wh.ID.String(),
			Address: wh.Address,
		}
		responseWarehouses = append(responseWarehouses, responseWarehouse)
	}

	log.WithField("count", len(responseWarehouses)).Debug("Parsed warehouses successfully")

	return responseWarehouses
}
