package di

import (
	"database/sql"

	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/delivery/controller/clientcontroller"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/clientrespository/clientrepositoryimpl"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/usecase/clientusecase"
)

func ConfigClientDI(conn *sql.DB) *clientcontroller.ClientController {
	clientRepository := clientrepositoryimpl.NewClientRepository(conn)
	clientUseCase := clientusecase.NewClientUseCase(clientRepository)
	clientController := clientcontroller.NewClientController(clientUseCase)

	return clientController

}
