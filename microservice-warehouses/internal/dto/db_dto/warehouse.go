package db_dto

import "github.com/google/uuid"

type WarehouseDTO struct {
	ID      uuid.UUID
	Address string
}

type CreateWarehouseDTO struct {
	Address string
}
