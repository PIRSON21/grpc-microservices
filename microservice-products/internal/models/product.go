package models

import (
	"errors"

	custErr "github.com/PIRSON21/grpc-microservices/microservice-products/pkg/errors"
	"github.com/google/uuid"
)

// Product represents a product entity in the system.
type Product struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_default_uuid();primary_key"` // UUID primary key
	Name        string         `gorm:"type:varchar(100);not null"`                       // Name of the product
	Description string         `gorm:"type:text;not null"`                               // Description of the product
	Params      map[string]any `gorm:"type:jsonb;default:'{}'"`                          // Additional parameters stored as JSONB
	Weight      float64        `gorm:"type:float;not null"`                              // Weight of the product
	BarCode     string         `gorm:"type:varchar(50);unique;not null"`                 // Unique barcode
}

// Validation errors
var (
	// ErrInvalidProductName indicates that the product name is invalid (empty or taller than 100 symbols).
	ErrInvalidProductName = errors.New("invalid product name")
	// ErrInvalidProductWeight indicates that the product weight is invalid (negative or zero).
	ErrInvalidProductWeight = errors.New("invalid product weight")
	// ErrInvalidProductBarCode indicates that the product barcode is invalid (empty or taller than 50 symbols).
	ErrInvalidProductBarCode = errors.New("invalid product bar code")
)

func (p *Product) Validate() []*custErr.ValidationError {
	var validationErrors []*custErr.ValidationError
	if p.Name == "" || len(p.Name) > 100 {
		validationErrors = append(validationErrors, &custErr.ValidationError{
			Field: "name",
			Err:   ErrInvalidProductName,
		})
	}
	if p.Weight <= 0 {
		validationErrors = append(validationErrors, &custErr.ValidationError{
			Field: "weight",
			Err:   ErrInvalidProductWeight,
		})
	}
	if p.BarCode == "" || len(p.BarCode) > 50 {
		validationErrors = append(validationErrors, &custErr.ValidationError{
			Field: "bar_code",
			Err:   ErrInvalidProductBarCode,
		})
	}
	return validationErrors
}
