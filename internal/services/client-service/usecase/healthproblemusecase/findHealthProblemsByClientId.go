package healthproblemusecase

import (
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

func (hpuc *HealthProblemUseCase) FindHealthProblemsByClientId(clientId uuid.UUID) ([]*entities.HealthProblem, error) {

	healthProblems, err := hpuc.healthProblemRepository.FindHealthProblemsByClientId(clientId)

	if err != nil {
		return nil, err
	}

	if healthProblems == nil {
		return nil, nil
	}

	return healthProblems, nil

}
