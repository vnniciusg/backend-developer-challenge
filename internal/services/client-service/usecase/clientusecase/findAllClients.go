package clientusecase

import "github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"

func (cuc *ClientUseCase) FindAllClients() ([]*entities.Client, error) {
	clients, err := cuc.clientRepository.FindAllClients()

	if err != nil {
		return nil, err
	}

	if clients == nil {
		return nil, nil
	}

	return clients, nil
}
