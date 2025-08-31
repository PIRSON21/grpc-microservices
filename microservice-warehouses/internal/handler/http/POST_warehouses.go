package http

import (
	"net/http"

	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateWarehouse handles POST requests to create a new warehouse.
func (h *WarehouseHTTPHandler) CreateWarehouse(c *gin.Context) {
	log := logrus.WithFields(logrus.Fields{
		"op":     "handler.WarehouseHTTPHandler.CreateWarehouse",
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	})

	log.Debug("Received request to create warehouse")

	var req dto.CreateWarehouseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.service.CreateWarehouse(c.Request.Context(), &req)
	if err != nil {
		log.Errorf("Error creating warehouse: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	log.Debugf("Warehouse created successfully: %+v", req)

	c.Status(http.StatusCreated)
}
