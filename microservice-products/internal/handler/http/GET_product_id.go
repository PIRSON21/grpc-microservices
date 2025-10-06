package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetProductByID handles GET requests to retrieve a product by its ID.
func (h *ProductHTTPHandler) GetProductByID(c *gin.Context) {
	log := logrus.WithFields(logrus.Fields{
		"op":     "handler.ProductHTTPHandler.GetProductByID",
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	})

	log.Debug("Received request to get product by ID")

	id := c.Param("id")
	if id == "" {
		log.Error("Product ID is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	product, err := h.service.GetProducts(c.Request.Context(), id)
	if err != nil {
		log.Errorf("Error retrieving product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(product) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product[0])
}
