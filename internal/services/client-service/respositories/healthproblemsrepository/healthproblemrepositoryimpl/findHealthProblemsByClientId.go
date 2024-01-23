package healthproblemrepositoryimpl

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/sqlstatements/healthproblemsqlstatements"
)

func (hpr *HealthProblemRepository) FindHealthProblemsByClientId(clientId uuid.UUID) ([]*entities.HealthProblems, error) {

	row, err := hpr.DB.Query(healthproblemsqlstatements.SelectHealthProblemByClientId, clientId)

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
			return nil, err
		}

		healthProblems = append(healthProblems, healthProblem)
	}

	return healthProblems, nil

}
