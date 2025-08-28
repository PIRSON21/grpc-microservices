package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// GetWarehouseByID handles GET requests to retrieve a warehouse by its ID.
func (h *WarehouseHandler) GetWarehouseByID(c *gin.Context) {
	log := logrus.WithFields(logrus.Fields{
		"op":     "handler.WarehouseHandler.GetWarehouseByID",
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	})

	log.Debug("Received request to get warehouses by ID")

	id, err := parseIDParam(c, "id")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	warehouses, err := h.service.GetWarehouses(c.Request.Context(), id)
	if err != nil {
		log.Errorf("Error retrieving warehouses with ID %s: %v", id, err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(warehouses) == 0 {
		log.Warnf("Warehouse with ID %s not found", id)
		c.JSON(404, gin.H{"error": "Warehouse not found"})
		return
	}

	c.JSON(200, warehouses)
}

func parseIDParam(c *gin.Context, param string) (string, error) {
	idStr := c.Param(param)
	if idStr == "" {
		return "", fmt.Errorf("missing %s parameter", param)
	}
	err := uuid.Validate(idStr)
	if err != nil {
		return "", fmt.Errorf("invalid %s parameter: %v", param, err)
	}

	return idStr, nil
}
