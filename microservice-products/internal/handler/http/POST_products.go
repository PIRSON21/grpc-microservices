package http

import (
	"errors"
	"net/http"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/dto"
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/service"
	"github.com/gin-gonic/gin"
)

// CreateProduct handles POST requests to create a new product.
func (h *ProductHTTPHandler) CreateProduct(c *gin.Context) {
	log := h.logger.With().Fields(map[string]any{
		"op":     "handler.ProductHTTPHandler.CreateProduct",
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	}).Logger()

	log.Debug().Msg("Received request to create product")

	var createRequest dto.CreateProductRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		log.Error().Msgf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createdProduct, err, validErr := h.service.CreateProduct(c.Request.Context(), &createRequest)
	if err != nil {
		if errors.Is(err, service.ValidationError) {
			log.Error().Msg("Product validation failed")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation Error", "details": validErr})
			return
		}
		log.Error().Msgf("Error creating product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}
