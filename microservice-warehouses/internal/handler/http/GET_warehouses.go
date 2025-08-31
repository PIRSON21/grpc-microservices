package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetWarehouses handles GET requests to retrieve a list of warehouses.
func (h *WarehouseHTTPHandler) GetWarehouses(c *gin.Context) {
	log := logrus.WithFields(logrus.Fields{
		"op":     "handler.WarehouseHTTPHandler.GetWarehouses",
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	})

	log.Debug("Received request to get warehouses")

	warehouses, err := h.service.GetWarehouses(c.Request.Context(), "")
	if err != nil {
		log.Errorf("Error retrieving warehouses: %v", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, warehouses)
}
