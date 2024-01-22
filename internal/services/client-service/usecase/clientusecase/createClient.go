package clientusecase

import (
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
)

func (cuc *ClientUseCase) CreateClient(client *request.CreateClientRequestDTO) error {
	err := cuc.clientRepository.CreateClient(client)

	if err != nil {
		return err
	}

	return nil
}
