package clientusecase

import (
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

func (cuc *ClientUseCase) FindClientById(id uuid.UUID) (*entities.Client, error) {

	client, err := cuc.clientRepository.FindClientById(id)

	if err != nil {
		return nil, err
	}

	if client == nil {
		return nil, nil
	}

	return client, nil

}
