package models

import "github.com/google/uuid"

// Product represents a product entity in the system.
type Product struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_default_uuid();primary_key"` // UUID primary key
	Name        string         `gorm:"type:varchar(100);not null"`                       // Name of the product
	Description string         `gorm:"type:text;not null"`                               // Description of the product
	Params      map[string]any `gorm:"type:jsonb;default:'{}'"`                          // Additional parameters stored as JSONB
	Weight      float64        `gorm:"type:float;not null"`                              // Weight of the product
	BarCode     string         `gorm:"type:varchar(50);unique;not null"`                 // Unique barcode
}
