package clientrespository

import (
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

type ClientRepository interface {
	FindAllClients() ([]*entities.Client, error)
	FindClientById(id uuid.UUID) (*entities.Client, error)
	CreateClient(client *request.CreateClientRequestDTO) (*entities.Client, error)
	UpdateClient(client *entities.Client) (*entities.Client, error)
	TopTenClientsWithHighestHealthRisk() ([]*entities.Client, error)
}
