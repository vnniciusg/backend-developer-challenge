package request

import "github.com/google/uuid"

type CreateHealthProblemRequestDTO struct {
	ClientId uuid.UUID `json:"client_id"`
	Name     string    `json:"name"`
	Grau     int       `json:"grau"`
}
