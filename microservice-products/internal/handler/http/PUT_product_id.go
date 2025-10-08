package http

import (
	"errors"
	"net/http"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/dto"
	custErr "github.com/PIRSON21/grpc-microservices/microservice-products/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// UpdateProduct manages the updating of an existing product by its ID.
func (h *ProductHTTPHandler) UpdateProduct(c *gin.Context) {
	log := h.logger.With().Fields(map[string]any{
		"op":     "handler.ProductHTTPHandler.UpdateProduct",
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	}).Logger()

	productID := c.Param("id")
	if productID == "" {
		log.Error().Msg("Product ID is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	var payload dto.UpdateProductRequest
	var err error
	if payload.ID, err = uuid.Parse(productID); err != nil {
		log.Error().Msgf("Invalid product ID format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			validErr := custErr.ProcessValidationErrors(ve)
			log.Error().Msg("Product validation failed")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation Error", "details": validErr})
			return
		}
		log.Error().Msgf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.service.UpdateProduct(c.Request.Context(), &payload); err != nil {
		// TODO: detailed error handling
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})

}
