package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josephjou806/go-jwt-claims-demo/internal/services"
)

type ClaimHandler struct {
	service services.ClaimService
}

func NewClaimHandler(s services.ClaimService) *ClaimHandler {
	return &ClaimHandler{service: s}
}

func (h *ClaimHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	claim, err := h.service.GetClaimByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "claim not found"})
		return
	}
	c.JSON(http.StatusOK, claim)
}
