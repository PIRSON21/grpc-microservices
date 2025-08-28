package nop

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/dto/db_dto"
	"github.com/sirupsen/logrus"
)

func (n *NOP) CreateWarehouse(_ context.Context, _ *db_dto.CreateWarehouseDTO) error {
	logrus.Info("NOP CreateWarehouse called - no operation performed")
	return nil
}
