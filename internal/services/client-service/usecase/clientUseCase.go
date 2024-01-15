package usecase

import "github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories"

type ClientUseCase struct {
	clientRepository *respositories.ClientRepository
}

func NewClientUseCase(clientRepository *respositories.ClientRepository) *ClientUseCase {
	return &ClientUseCase{clientRepository: clientRepository}
}
