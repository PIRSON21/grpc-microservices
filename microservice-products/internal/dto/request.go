package dto

// CreateProductRequest represents the payload for creating a new product.
type CreateProductRequest struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description"`
	Params      map[string]any `json:"params"`
	Weight      float64        `json:"weight" validate:"required,gt=0"`
	BarCode     string         `json:"bar_code" validate:"required"`
}
