package healthproblemrepository

import (
	"database/sql"

	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence/utils"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/healthproblemrepository/sqlstatments"
)

func (hpr *HealthProblemsRepository) CreateHealthProblem(healthProblem *request.CreateHealthProblemRequestDTO) (*entities.HealthProblems, error) {

	var healthProblemResponse *entities.HealthProblems

	err := utils.WithTransaction(hpr.DB, func(tx *sql.Tx) error {

		created_at, err := date.GetCurrentTime()

		if err != nil {
			return err
		}

		updated_at := created_at

		row, err := tx.Query(sqlstatments.InsertHealthProblem, healthProblem.ClientId, healthProblem.Name, healthProblem.Grau, created_at, updated_at)

		if err != nil {
			return err
		}

		defer row.Close()

		err = row.Scan(&healthProblemResponse.Id, &healthProblemResponse.ClientId, &healthProblemResponse.Name, &healthProblemResponse.Grau, &healthProblemResponse.CreatedAt, &healthProblemResponse.UpdatedAt)

		return nil

	})

	return healthProblemResponse, err
}
