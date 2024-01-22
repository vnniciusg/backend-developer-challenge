package clientrepositoryimpl

import (
	"database/sql"

	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/response"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/healthproblemsrepository/healthproblemrepositoryimpl"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/sqlstatements/clientsqlstatements"
)

func (cr *ClientRepository) FindAllClients() ([]*response.GetClientResponseDTO, error) {

	row, err := cr.DB.Query(clientsqlstatements.SelectAllClients)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	defer row.Close()

	var clients []*response.GetClientResponseDTO

	for row.Next() {
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
			return nil, err
		}

		healProblemRepository := healthproblemrepositoryimpl.NewHealthProblemRepository(cr.DB)

		healthProblem, err := healProblemRepository.FindHealthProblemsByClientId(client.Id)

		if err != nil {
			return nil, err
		}

		client.HealthProblems = healthProblem

		clients = append(clients, client)
	}

	return clients, nil

}
