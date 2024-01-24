package healthproblemusecase

import (
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
)

func (hpuc *HealthProblemUseCase) UpdateHealthProblem(id uuid.UUID, healthProblem *request.UpdateHealthProblemRequestDTO) error {
	err := hpuc.UpdateHealthProblem(id, healthProblem)

	if err != nil {
		return err
	}

	return nil
}
