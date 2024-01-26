package clientrepositoryimpl

import (
	"database/sql"
	"strconv"

	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence/utils"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	queryUtils "github.com/vnniciusg/backend-developer-challenge/utils"
)

func (cr *ClientRepository) UpdateClient(client *entities.Client) error {

	err := utils.WithTransaction(cr.DB, func(tx *sql.Tx) error {
		err := utils.CheckEntityExists(cr.DB, "SELECT id FROM tb_clients WHERE id = $1", client.Id, utils.Client)
		if err != nil {
			return err
		}

		query, args := buildClientUpdateQuery(client)

		_, err = tx.Exec(query, args...)

		if err != nil {
			return err
		}

		return nil

	})

	return err
}

func buildClientUpdateQuery(client *entities.Client) (string, []interface{}) {

	var query string
	var args []interface{}
	var argsCounts int

	updated_at, err := date.GetCurrentTime()

	if err != nil {
		return "", nil
	}

	query += "UPDATE tb_clients SET "

	queryUtils.QueryAppend(&query, &args, &argsCounts, "", "name=$"+strconv.Itoa(argsCounts+1), client.Name)
	queryUtils.QueryAppend(&query, &args, &argsCounts, ",", "birth_date=$"+strconv.Itoa(argsCounts+1), client.BirthDate)
	queryUtils.QueryAppend(&query, &args, &argsCounts, ",", "sexo=$"+strconv.Itoa(argsCounts+1), client.Sexo)
	queryUtils.QueryAppend(&query, &args, &argsCounts, ",", "updated_at=$"+strconv.Itoa(argsCounts+1), updated_at)
	queryUtils.QueryAppend(&query, &args, &argsCounts, " WHERE ", "id=$"+strconv.Itoa(argsCounts+1), client.Id)

	return query, args
}
