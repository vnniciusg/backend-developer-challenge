package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

type GetClientResponseDTO struct {
	Id             uuid.UUID                  `json:"id"`
	Name           string                     `json:"name"`
	BirthDate      time.Time                  `json:"birth_date"`
	Sexo           string                     `json:"sexo"`
	CreatedAt      time.Time                  `json:"created_at"`
	UpdatedAt      time.Time                  `json:"updated_at"`
	HealthProblems []*entities.HealthProblems `json:"health_problems"`
}
