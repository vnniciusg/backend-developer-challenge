package clientusecase

import (
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

func (cuc *ClientUseCase) CreateClient(client *entities.Client) error {
	err := cuc.clientRepository.CreateClient(client)

	if err != nil {
		return err
	}

	return nil
}
