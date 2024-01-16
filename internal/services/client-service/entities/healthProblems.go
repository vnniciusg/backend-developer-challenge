package entities

import (
	"time"

	"github.com/google/uuid"
)

type HealthProblems struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ClientId  string    `json:"client_id"`
	Grau      int       `json:"grau"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
