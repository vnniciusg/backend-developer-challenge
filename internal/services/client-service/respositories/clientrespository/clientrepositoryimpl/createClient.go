package clientrepositoryimpl

import (
	"database/sql"

	"time"

	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence/utils"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/clientrespository/clientrepositoryimpl/sqlstatements"
	healproblemssqlstatements "github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/healthproblemrepository/sqlstatements"
)

func (cr *ClientRepository) CreateClient(client *request.CreateClientRequestDTO) (*entities.Client, error) {

	var clientResponse *entities.Client

	err := utils.WithTransaction(cr.DB, func(tx *sql.Tx) error {

		created_at, err := date.GetCurrentTime()

		if err != nil {
			return err
		}

		updated_at := created_at

		layout := "02/01/2006"

		BirthDate, err := time.Parse(layout, client.BirthDate)

		if err != nil {
			return err
		}

		row := tx.QueryRow(sqlstatements.InsertClient, client.Name, BirthDate, client.Sexo, created_at, updated_at)

		clientResponse = &entities.Client{}

		err = row.Scan(&clientResponse.Id, &clientResponse.Name, &clientResponse.BirthDate, &clientResponse.Sexo, &clientResponse.CreatedAt, &clientResponse.UpdatedAt)
		if err != nil {
			return err
		}
		for _, healthProblem := range client.HealthProblems {

			healthProblemResponse := &entities.HealthProblems{}

			row := tx.QueryRow(healproblemssqlstatements.InsertHealthProblem, clientResponse.Id, healthProblem.Name, healthProblem.Grau, created_at, updated_at)

			err = row.Scan(&healthProblemResponse.Id, &healthProblemResponse.ClientId, &healthProblemResponse.Name, &healthProblemResponse.Grau, &healthProblemResponse.CreatedAt, &healthProblemResponse.UpdatedAt)

			if err != nil {
				return err
			}

			clientResponse.HealthProblems = append(clientResponse.HealthProblems, healthProblemResponse)

		}

		return nil

	})

	return clientResponse, err
}
