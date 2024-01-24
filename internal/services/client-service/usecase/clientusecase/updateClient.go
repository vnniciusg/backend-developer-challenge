package clientusecase

import (
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
)

func (cuc *ClientUseCase) UpdateClient(id uuid.UUID, client *request.UpdateClientRequestDTO) error {
	err := cuc.UpdateClient(id, client)

	if err != nil {
		return err
	}

	return nil
}
