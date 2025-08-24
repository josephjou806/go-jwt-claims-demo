package services

import (
	"context"

	"github.com/josephjou806/go-jwt-claims-demo/internal/models"
	"github.com/josephjou806/go-jwt-claims-demo/internal/repository"
)

type ClaimService interface {
	GetClaimByID(ctx context.Context, id string) (*models.Claim, error)
}

type claimService struct {
	repo repository.ClaimRepository
}

func NewClaimService(repo repository.ClaimRepository) ClaimService {
	return &claimService{repo: repo}
}

func (s *claimService) GetClaimByID(ctx context.Context, id string) (*models.Claim, error) {
	return s.repo.FindByID(id)
}
