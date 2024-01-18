package clientcontroller

import "github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/usecase/clientusecase"

type ClientController struct {
	clientUseCase *clientusecase.ClientUseCase
}

func NewClientController(clientUseCase *clientusecase.ClientUseCase) *ClientController {
	return &ClientController{clientUseCase: clientUseCase}
}
