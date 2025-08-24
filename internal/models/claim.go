package models

import "time"

type Claim struct {
	ID       string    `json:"id"`
	MemberID string    `json:"memberId"`
	NDC      string    `json:"ndc"`
	Amount   float64   `json:"amount"`
	Status   string    `json:"status"`
	FillDate time.Time `json:"fillDate"`
}
