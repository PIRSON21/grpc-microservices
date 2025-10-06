package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteProduct handles DELETE requests to delete a product by its ID.
func (h *ProductHTTPHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Status(http.StatusNoContent)
}
