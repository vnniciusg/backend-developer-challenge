package clientrepositoryimpl

import (
	"database/sql"

	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence/utils"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/clientrespository/clientrepositoryimpl/sqlstatements"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/healthproblemrepository"
)

func (cr *ClientRepository) CreateClient(client *request.CreateClientRequestDTO) (*entities.Client, error) {

	var clientResponse *entities.Client

	err := utils.WithTransaction(cr.DB, func(tx *sql.Tx) error {

		healthProblemRepository := healthproblemrepository.NewHealthProblemsRepository(cr.DB)

		created_at, err := date.GetCurrentTime()

		if err != nil {
			return err
		}

		updated_at := created_at

		row, err := tx.Query(sqlstatements.InsertClient, client.Name, client.BirthDate, client.Sexo, created_at, updated_at)

		if err != nil {
			return err
		}

		defer row.Close()

		err = row.Scan(&clientResponse.Id, &clientResponse.Name, &clientResponse.BirthDate, &clientResponse.Sexo, &clientResponse.Sexo, &clientResponse.CreatedAt, &clientResponse.UpdatedAt)

		if err != nil {
			return err
		}

		for _, healthProblem := range client.HealthProblems {
			healthProblem.ClientId = clientResponse.Id

			healthProblemResponse, err := healthProblemRepository.CreateHealthProblem(healthProblem)

			if err != nil {
				return err
			}

			clientResponse.HealthProblems = append(clientResponse.HealthProblems, healthProblemResponse)

		}

		return nil

	})

	return clientResponse, err
}
