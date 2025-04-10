package types

import "github.com/google/uuid"

type Benefit struct {
	BenefitID   uuid.UUID `json:"benefit_id"`
	Name        string    `json:"benefit_name"`
	Description string    `json:"description"`
	Amount      float32   `json:"amount"`
	Frequency   int       `json:"frequency"`
}
