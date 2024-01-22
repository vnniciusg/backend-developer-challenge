package clientrepositoryimpl

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/sqlstatements/clientsqlstatements"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/sqlstatements/healthproblemsqlstatements"
)

func (cr *ClientRepository) FindClientById(id uuid.UUID) (*entities.Client, error) {

	row := cr.DB.QueryRow(clientsqlstatements.SelectClientById, id)

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
		if err == sql.ErrNoRows {
			return nil, nil

		}
		return nil, err
	}

	rows, err := cr.DB.Query(healthproblemsqlstatements.SelectHealthProblemById, client.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	defer rows.Close()

	var healthProblems []*entities.HealthProblems
	for rows.Next() {
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

	return client, nil

}
