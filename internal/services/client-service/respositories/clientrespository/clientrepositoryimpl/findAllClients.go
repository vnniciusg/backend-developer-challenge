package clientrepositoryimpl

import (
	"database/sql"

	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/sqlstatements/clientsqlstatements"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/sqlstatements/healthproblemsqlstatements"
)

func (cr *ClientRepository) FindAllClients() ([]*entities.Client, error) {

	row, err := cr.DB.Query(clientsqlstatements.SelectAllClients)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	defer row.Close()

	var clients []*entities.Client

	for row.Next() {
		client := &entities.Client{}

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

		row, err := cr.DB.Query(healthproblemsqlstatements.SelectHealthProblemById, client.Id)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}

		defer row.Close()

		var healthProblems []*entities.HealthProblems

		for row.Next() {
			healthProblem := &entities.HealthProblems{}

			err := row.Scan(
				&healthProblem.Id,
				&healthProblem.Name,
				&healthProblem.ClientId,
				&healthProblem.Grau,
				&healthProblem.CreatedAt,
				&healthProblem.UpdatedAt,
			)

			if err != nil {
				if err == sql.ErrNoRows {
					return nil, nil
				}
				return nil, err
			}

			healthProblems = append(healthProblems, healthProblem)

		}

		client.HealthProblems = healthProblems
		clients = append(clients, client)
	}

	return clients, nil

}
