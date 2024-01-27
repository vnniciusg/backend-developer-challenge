package entities

import (
	"time"

	"github.com/google/uuid"
)

type HealthProblem struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ClientId  uuid.UUID `json:"client_id"`
	Grau      int       `json:"grau"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewHealthProblem(name string, clientId uuid.UUID, grau int) *HealthProblem {
	return &HealthProblem{
		Name:     name,
		ClientId: clientId,
		Grau:     grau,
	}
}
