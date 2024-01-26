package healthproblemrepositoryimpl

import (
	"database/sql"

	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence/utils"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/sqlstatements/healthproblemsqlstatements"
)

func (hpr *HealthProblemRepository) CreateHealthProblem(healthProblems []*entities.HealthProblems) error {

	err := utils.WithTransaction(hpr.DB, func(tx *sql.Tx) error {

		for _, healthProblem := range healthProblems {

			err := createHealthProblem(tx, healthProblem)

			if err != nil {
				return err
			}
		}

		return nil

	})

	return err

}

func createHealthProblem(tx *sql.Tx, healthProblem *entities.HealthProblems) error {

	created_at, err := date.GetCurrentTime()

	if err != nil {
		return err
	}

	updated_at := created_at

	_, err = tx.Exec(healthproblemsqlstatements.InsertHealthProblem, healthProblem.Id, healthProblem.Name, healthProblem.Grau, created_at, updated_at)

	if err != nil {
		return err
	}

	return nil
}
