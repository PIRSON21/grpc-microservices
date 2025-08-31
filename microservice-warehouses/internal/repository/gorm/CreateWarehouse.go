package gorm

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/dto/db_dto"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/models"
	"github.com/sirupsen/logrus"
)

func (gorm *DbGorm) CreateWarehouse(ctx context.Context, warehouse *db_dto.CreateWarehouseDTO) error {
	log := logrus.WithFields(logrus.Fields{
		"op":      "repository.DbGorm.CreateWarehouse",
		"address": warehouse.Address,
	})

	model := models.Warehouse{
		Address: warehouse.Address,
	}

	log.Debug("Creating a new warehouse in the database")

	if err := gorm.db.WithContext(ctx).Model(&models.Warehouse{}).Create(&model).Error; err != nil {
		log.Error("Error creating warehouse: ", err)
		return err
	}

	log.Debug("Warehouse created successfully")

	return nil
}
