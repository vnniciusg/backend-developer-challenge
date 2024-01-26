package clientrespository

import (
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/response"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

type ClientRepository interface {
	FindAllClients() ([]*response.GetClientResponseDTO, error)
	FindClientById(id uuid.UUID) (*response.GetClientResponseDTO, error)
	CreateClient(client *entities.Client) error
	UpdateClient(client *entities.Client) error
	// TopTenClientsWithHighestHealthRisk() ([]*entities.Client, error)
}
