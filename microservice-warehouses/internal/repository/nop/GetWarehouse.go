package nop

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/dto/db_dto"
	"github.com/google/uuid"
)

func (n *NOP) GetWarehouses(_ context.Context, id string) ([]*db_dto.WarehouseDTO, error) {
	if id == "" {
		return returnAllWarehouses()
	}

	return returnSingleWarehouse(id)

}

func returnAllWarehouses() ([]*db_dto.WarehouseDTO, error) {
	return []*db_dto.WarehouseDTO{
		{
			ID:      uuid.New(),
			Address: "New York",
		},
		{
			ID:      uuid.New(),
			Address: "Los Angeles",
		},
	}, nil
}

func returnSingleWarehouse(id string) ([]*db_dto.WarehouseDTO, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return []*db_dto.WarehouseDTO{
		{
			ID:      uid,
			Address: "New York",
		},
	}, nil
}
