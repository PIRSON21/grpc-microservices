package repository

import (
	"context"

	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/dto/db_dto"
)

type GetWarehouses interface {
	GetWarehouses(context.Context, string) ([]*db_dto.WarehouseDTO, error)
}

type PostWarehouses interface {
	CreateWarehouse(context.Context, *db_dto.CreateWarehouseDTO) error
}

type WarehouseRepository interface {
	GetWarehouses
	PostWarehouses
}
