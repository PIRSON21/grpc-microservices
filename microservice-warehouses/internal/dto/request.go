package dto

type CreateWarehouseRequest struct {
	Address string `json:"address" binding:"required"`
}
