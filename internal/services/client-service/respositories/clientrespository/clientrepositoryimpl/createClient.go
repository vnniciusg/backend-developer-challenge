package clientrepositoryimpl

import (
	"database/sql"
	"time"

	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence/utils"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/sqlstatements/clientsqlstatements"
)

func (cr *ClientRepository) CreateClient(client *entities.Client) error {

	err := utils.WithTransaction(cr.DB, func(tx *sql.Tx) error {

		err := entities.ValidateClient(client)

		if err != nil {
			return err
		}

		err = insertClient(tx, client.Name, client.BirthDate, client.Sexo)

		if err != nil {
			return err
		}

		return nil

	})

	return err
}

func insertClient(tx *sql.Tx, name string, birthDate time.Time, sexo string) error {

	created_at, err := date.GetCurrentTime()

	if err != nil {
		return err
	}

	updated_at := created_at

	_, err = tx.Exec(clientsqlstatements.InsertClient, name, birthDate, sexo, created_at, updated_at)

	if err != nil {
		return err
	}

	return nil
}
