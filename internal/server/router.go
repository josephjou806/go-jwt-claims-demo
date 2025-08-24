package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josephjou806/go-jwt-claims-demo/internal/config"
	"github.com/josephjou806/go-jwt-claims-demo/internal/handlers"
	"github.com/josephjou806/go-jwt-claims-demo/internal/middleware"
	"github.com/josephjou806/go-jwt-claims-demo/internal/services"
	"github.com/josephjou806/go-jwt-claims-demo/internal/token"
)

func NewRouter(cfg config.Config, svc services.ClaimService, tm token.TokenManager) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.NewCORS(cfg.CORSAllowedOrigin).Handler())

	authH := handlers.NewAuthHandler(tm)
	claimH := handlers.NewClaimHandler(svc)
	jwtMw := middleware.NewJWTAuth(tm)

	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	r.POST("/login", authH.Login)

	api := r.Group("/claims")
	api.Use(jwtMw.Handler())
	{
		api.GET(":id", claimH.GetByID)
	}
	return r
}
