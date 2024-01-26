package healthproblemusecase

import (
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

func (hpuc *HealthProblemUseCase) CreateHealthProblem(healthProblems []*entities.HealthProblems) error {

	err := hpuc.healthProblemRepository.CreateHealthProblem(healthProblems)

	if err != nil {
		return err
	}

	return nil

}
