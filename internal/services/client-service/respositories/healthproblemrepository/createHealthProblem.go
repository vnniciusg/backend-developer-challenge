package healthproblemrepository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence/utils"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/healthproblemrepository/sqlstatements"
)

func (hpr *HealthProblemsRepository) CreateHealthProblem(clientId uuid.UUID, healthProblem *request.CreateHealthProblemRequestDTO) (*entities.HealthProblems, error) {

	var healthProblemResponse *entities.HealthProblems

	err := utils.WithTransaction(hpr.DB, func(tx *sql.Tx) error {

		healthProblemResponse := &entities.HealthProblems{}

		created_at, err := date.GetCurrentTime()

		if err != nil {
			return err
		}

		updated_at := created_at

		row := tx.QueryRow(sqlstatements.InsertHealthProblem, clientId, &healthProblem.Name, &healthProblem.Grau, created_at, updated_at)

		err = row.Scan(&healthProblemResponse.Id, &healthProblemResponse.ClientId, &healthProblemResponse.Name, &healthProblemResponse.Grau, &healthProblemResponse.CreatedAt, &healthProblemResponse.UpdatedAt)
		if err != nil {
			return err
		}

		return nil

	})

	return healthProblemResponse, err
}
