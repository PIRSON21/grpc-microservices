package gorm

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/dto/db_dto"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/models"
	"github.com/sirupsen/logrus"
)

// GetWarehouses retrieves warehouses from the database.
// If uuid is provided, it filters the warehouses by the given UUID.
//
// It returns a slice of WarehouseDTO and an error if any occurs during the operation.
func (gorm *DbGorm) GetWarehouses(ctx context.Context, uuid string) ([]*db_dto.WarehouseDTO, error) {
	log := logrus.WithFields(logrus.Fields{
		"op":   "repository.DbGorm.GetWarehouses",
		"uuid": uuid,
	})

	log.Debug("Fetching warehouses from the database")

	var warehouses []models.Warehouse
	query := gorm.db.WithContext(ctx).Model(&models.Warehouse{})
	if uuid != "" {
		query = query.Where("warehouse_id = ?", uuid)
	}
	if err := query.Find(&warehouses).Error; err != nil {
		log.Error("Error fetching warehouses: ", err)
		return nil, err
	}
	result := make([]*db_dto.WarehouseDTO, len(warehouses))
	for i := range warehouses {
		result[i] = &db_dto.WarehouseDTO{
			ID:      warehouses[i].ID,
			Address: warehouses[i].Address,
		}
	}

	log.Debugf("Fetched %d warehouses", len(result))

	return result, nil
}
