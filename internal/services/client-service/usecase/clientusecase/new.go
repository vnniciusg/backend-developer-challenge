package clientusecase

import "github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/clientrespository"

type ClientUseCase struct {
	clientRepository clientrespository.ClientRepository
}

func NewClientUseCase(clientRepository clientrespository.ClientRepository) *ClientUseCase {
	return &ClientUseCase{clientRepository: clientRepository}
}
