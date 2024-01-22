package clientrepositoryimpl

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/response"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/healthproblemsrepository/healthproblemrepositoryimpl"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/sqlstatements/clientsqlstatements"
)

func (cr *ClientRepository) FindClientById(id uuid.UUID) (*response.GetClientResponseDTO, error) {

	row := cr.DB.QueryRow(clientsqlstatements.SelectClientById, id)

	client := &response.GetClientResponseDTO{}

	err := row.Scan(
		&client.Id,
		&client.Name,
		&client.BirthDate,
		&client.Sexo,
		&client.CreatedAt,
		&client.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil

		}
		return nil, err
	}

	healProblemRepository := healthproblemrepositoryimpl.NewHealthProblemRepository(cr.DB)

	healthProblem, err := healProblemRepository.FindHealthProblemsByClientId(client.Id)

	if err != nil {
		return nil, err
	}

	client.HealthProblems = healthProblem

	return client, nil

}
