package clientusecase

import (
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

func (cuc *ClientUseCase) CreateClient(client *request.CreateClientRequestDTO) (*entities.Client, error) {
	clientResponse, err := cuc.clientRepository.CreateClient(client)

	if err != nil {
		return nil, err
	}

	return clientResponse, nil
}
