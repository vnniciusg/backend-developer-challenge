package clientusecase

import "github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/response"

func (cuc *ClientUseCase) FindAllClients() ([]*response.GetClientResponseDTO, error) {
	clients, err := cuc.clientRepository.FindAllClients()

	if err != nil {
		return nil, err
	}

	if clients == nil {
		return nil, nil
	}

	return clients, nil
}
