package repository

import (
	"errors"
	"time"

	"github.com/josephjou806/go-jwt-claims-demo/internal/models"
)

type ClaimRepository interface {
	FindByID(id string) (*models.Claim, error)
}

type inMemoryClaimRepository struct {
	data map[string]models.Claim
}

func NewInMemoryClaimRepository() ClaimRepository {
	return &inMemoryClaimRepository{data: seed()}
}

func (r *inMemoryClaimRepository) FindByID(id string) (*models.Claim, error) {
	c, ok := r.data[id]
	if !ok {
		return nil, errors.New("claim not found")
	}
	return &c, nil
}

func seed() map[string]models.Claim {
	return map[string]models.Claim{
		"1001": {ID: "1001", MemberID: "M-001", NDC: "0002-8215-01", Amount: 15.75, Status: "PAID", FillDate: time.Now().AddDate(0, 0, -5)},
		"1002": {ID: "1002", MemberID: "M-002", NDC: "0007-3163-13", Amount: 3.99, Status: "REJECTED", FillDate: time.Now().AddDate(0, 0, -2)},
	}
}
