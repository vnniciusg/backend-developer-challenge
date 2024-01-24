package healthproblemrepositoryimpl

import (
	"database/sql"
	"strconv"

	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence/utils"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	queryUtils "github.com/vnniciusg/backend-developer-challenge/utils"
)

func (hpr *HealthProblemRepository) UpdateHealthProblem(id uuid.UUID, healthProblem []request.UpdateHealthProblemRequestDTO) error {

	err := utils.WithTransaction(hpr.DB, func(tx *sql.Tx) error {
		err := utils.CheckEntityExists(hpr.DB, "SELECT id FROM tb_health_problems WHERE id = $1", id, utils.HealthProblem)
		if err != nil {
			return err
		}

		for _, hp := range healthProblem {

			query, args := buildHealthProblemUpdateQuery(id, hp)

			_, err := tx.Exec(query, args...)

			if err != nil {
				return err
			}

			return nil

		}

		return nil
	})

	return err

}

func buildHealthProblemUpdateQuery(id uuid.UUID, healthProblem request.UpdateHealthProblemRequestDTO) (string, []interface{}) {

	var query string
	var args []interface{}
	var argsCounts int

	updated_at, err := date.GetCurrentTime()

	if err != nil {
		return "", nil
	}

	query += "UPDATE tb_health_problems SET "
	queryUtils.QueryAppend(&query, &args, &argsCounts, "", "name=$"+strconv.Itoa(argsCounts+1), healthProblem.Name)
	queryUtils.QueryAppend(&query, &args, &argsCounts, ",", "grau=$"+strconv.Itoa(argsCounts+1), healthProblem.Grau)
	queryUtils.QueryAppend(&query, &args, &argsCounts, ",", "updated_at=$"+strconv.Itoa(argsCounts+1), updated_at)
	queryUtils.QueryAppend(&query, &args, &argsCounts, " WHERE ", "id=$"+strconv.Itoa(argsCounts+1), id)

	return query, args
}
