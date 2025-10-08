package dto

import "github.com/google/uuid"

// CreateProductRequest represents the payload for creating a new product.
type CreateProductRequest struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description"`
	Params      map[string]any `json:"params"`
	Weight      float64        `json:"weight" validate:"required,gt=0"`
	BarCode     string         `json:"bar_code" validate:"required"`
}

// UpdateProductRequest represents the payload for updating an existing product.
type UpdateProductRequest struct {
	ID          uuid.UUID      `json:"-"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Params      map[string]any `json:"params"`
	Weight      float64        `json:"weight" binding:"gt=0"`
	BarCode     string         `json:"bar_code"`
}
