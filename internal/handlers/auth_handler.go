package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/josephjou806/go-jwt-claims-demo/internal/token"
)

type AuthHandler struct {
	token token.TokenManager
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResp struct {
	Token string `json:"token"`
}

func NewAuthHandler(tm token.TokenManager) *AuthHandler {
	return &AuthHandler{token: tm}
}

// Demo-only: accepts any non-empty username/password and issues a JWT
func (h *AuthHandler) Login(c *gin.Context) {
	var req loginReq
	if err := c.BindJSON(&req); err != nil || req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}
	t, err := h.token.Generate(req.Username, 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to issue token"})
		return
	}
	c.JSON(http.StatusOK, loginResp{Token: t})
}
