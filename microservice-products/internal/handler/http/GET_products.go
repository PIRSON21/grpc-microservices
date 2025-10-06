package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts handles GET requests to retrieve a list of products.
func (h *ProductHTTPHandler) GetProducts(c *gin.Context) {
	log := h.logger.With().Fields(map[string]any{
		"op":     "handler.ProductHTTPHandler.GetProducts",
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	}).Logger()

	log.Debug().Msg("Received request to get products")

	products, err := h.service.GetProducts(c.Request.Context(), "")
	if err != nil {
		log.Error().Msgf("Error retrieving products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, products)
}
