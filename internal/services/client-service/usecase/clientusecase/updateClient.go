package clientusecase

import (
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

func (cuc *ClientUseCase) UpdateClient(client *entities.Client) error {
	err := cuc.UpdateClient(client)

	if err != nil {
		return err
	}

	return nil
}
