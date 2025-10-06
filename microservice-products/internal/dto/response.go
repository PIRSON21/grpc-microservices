package dto

// GetProductResponse represents the response structure for retrieving a product.
type GetProductResponse struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Params      map[string]any `json:"params"`
	Weight      float64        `json:"weight"`
	BarCode     string         `json:"bar_code"`
}

// CreateProductResponse represents the response after creating a new product.
type CreateProductResponse struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Params      map[string]any `json:"params"`
	Weight      float64        `json:"weight"`
	BarCode     string         `json:"bar_code"`
}
