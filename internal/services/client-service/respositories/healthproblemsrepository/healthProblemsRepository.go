package healthproblemsrepository

import (
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

type HealthProblemsRepository interface {
	FindHealthProblemsByClientId(clientId uuid.UUID) ([]*entities.HealthProblems, error)
	CreateHealthProblem(client []*entities.HealthProblems) error
	UpdateHealthProblem(client *entities.HealthProblems) error
}
