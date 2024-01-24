package clientrespository

import (
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/response"
)

type ClientRepository interface {
	FindAllClients() ([]*response.GetClientResponseDTO, error)
	FindClientById(id uuid.UUID) (*response.GetClientResponseDTO, error)
	CreateClient(client *request.CreateClientRequestDTO) error
	UpdateClient(id uuid.UUID, client *request.UpdateClientRequestDTO) error
	// TopTenClientsWithHighestHealthRisk() ([]*entities.Client, error)
}
