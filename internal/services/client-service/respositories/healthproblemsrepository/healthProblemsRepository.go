package healthproblemsrepository

import (
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

type HealthProblemsRepository interface {
	FindHealthProblemsByClientId(clientId uuid.UUID) ([]*entities.HealthProblems, error)
	CreateHealthProblem(clientId uuid.UUID, healthProblem []request.CreateHealthProblemRequestDTO) error
	UpdateHealthProblem(id uuid.UUID, healthProblem request.UpdateHealthProblemRequestDTO) error
}
