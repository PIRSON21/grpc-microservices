package http

import (
	"net/http"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateProduct handles POST requests to create a new product.
func (h *ProductHTTPHandler) CreateProduct(c *gin.Context) {
	log := logrus.WithFields(logrus.Fields{
		"op":     "handler.ProductHTTPHandler.CreateProduct",
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	})

	log.Debug("Received request to create product")

	var createRequest dto.CreateProductRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		log.Errorf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createdProduct, err := h.service.CreateProduct(c.Request.Context(), &createRequest)
	if err != nil {
		log.Errorf("Error creating product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}
