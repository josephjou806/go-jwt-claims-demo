package main

import (
	"log"
	"net/http"

	"github.com/josephjou806/go-jwt-claims-demo/internal/config"
	"github.com/josephjou806/go-jwt-claims-demo/internal/repository"
	"github.com/josephjou806/go-jwt-claims-demo/internal/server"
	"github.com/josephjou806/go-jwt-claims-demo/internal/services"
	"github.com/josephjou806/go-jwt-claims-demo/internal/token"
)

func main() {
	cfg := config.Load()

	// DI wiring
	repo := repository.NewInMemoryClaimRepository()
	service := services.NewClaimService(repo)
	tm := token.NewTokenManager(cfg.JWTSecret)

	r := server.NewRouter(cfg, service, tm)

	log.Printf("starting server on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
