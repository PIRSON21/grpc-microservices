package models

import "github.com/google/uuid"

type Warehouse struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:warehouse_id"`
	Address string    `gorm:"type:varchar(255);not null;column:warehouse_address"`
}
