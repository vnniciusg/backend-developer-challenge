package healthproblemsrepository

import (
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

type HealthProblemsRepository interface {
	FindHealthProblemsByClientId(clientId uuid.UUID) ([]*entities.HealthProblem, error)
	CreateHealthProblem(client []*entities.HealthProblem) error
	UpdateHealthProblem(client *entities.HealthProblem) error
}
