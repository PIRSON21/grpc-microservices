package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetProducts handles GET requests to retrieve a list of products.
func (h *ProductHTTPHandler) GetProducts(c *gin.Context) {
	log := logrus.WithFields(logrus.Fields{
		"op":     "handler.ProductHTTPHandler.GetProducts",
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	})

	log.Debug("Received request to get products")

	products, err := h.service.GetProducts(c.Request.Context(), "")
	if err != nil {
		log.Errorf("Error retrieving products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, products)
}
