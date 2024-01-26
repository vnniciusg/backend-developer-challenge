package healthproblemusecase

import (
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

func (hpuc *HealthProblemUseCase) UpdateHealthProblem(healthProblem *entities.HealthProblems) error {
	err := hpuc.UpdateHealthProblem(healthProblem)

	if err != nil {
		return err
	}

	return nil
}
