package healthproblemusecase

import (
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
)

func (hpuc *HealthProblemUseCase) CreateHealthProblem(clientId uuid.UUID, healthProblem []*request.CreateHealthProblemRequestDTO) error {

	err := hpuc.healthProblemRepository.CreateHealthProblem(clientId, healthProblem)

	if err != nil {
		return err
	}

	return nil

}
