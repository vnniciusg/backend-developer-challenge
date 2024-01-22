package clientrepositoryimpl

import (
	"database/sql"
	"errors"

	"time"

	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence/utils"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/sqlstatements/clientsqlstatements"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/sqlstatements/healthproblemsqlstatements"
)

func (cr *ClientRepository) CreateClient(client *request.CreateClientRequestDTO) (*entities.Client, error) {

	var createdClient *entities.Client

	err := utils.WithTransaction(cr.DB, func(tx *sql.Tx) error {

		created_at, err := date.GetCurrentTime()

		if err != nil {
			return err
		}

		updated_at := created_at

		birthDate, err := date.ParseDate(client.BirthDate)

		if err != nil {
			return errors.New("Falha ao converter data de nascimento")
		}

		createdClient, err = insertClient(tx, client.Name, birthDate, client.Sexo, created_at, updated_at)

		if err != nil {
			return err
		}

		for _, healthProblem := range client.HealthProblems {
			err = insertHealthProblem(tx, createdClient, healthProblem, created_at, updated_at)
			if err != nil {
				return err
			}

		}

		return nil

	})

	return createdClient, err
}

func insertClient(tx *sql.Tx, name string, birthDate time.Time, sexo string, createdAt time.Time, updatedAt time.Time) (*entities.Client, error) {
	row := tx.QueryRow(clientsqlstatements.InsertClient, name, birthDate, sexo, createdAt, updatedAt)

	createdClient := &entities.Client{}

	if err := row.Scan(&createdClient.Id, &createdClient.Name, &createdClient.BirthDate, &createdClient.Sexo, &createdClient.CreatedAt, &createdClient.UpdatedAt); err != nil {
		return nil, err
	}

	return createdClient, nil
}

func insertHealthProblem(tx *sql.Tx, client *entities.Client, healthProblem *request.CreateHealthProblemRequestDTO, createdAt time.Time, updatedAt time.Time) error {
	healthProblemResponse := &entities.HealthProblems{}

	row := tx.QueryRow(healthproblemsqlstatements.InsertHealthProblem, client.Id, healthProblem.Name, healthProblem.Grau, createdAt, updatedAt)

	if err := row.Scan(&healthProblemResponse.Id, &healthProblemResponse.ClientId, &healthProblemResponse.Name, &healthProblemResponse.Grau, &healthProblemResponse.CreatedAt, &healthProblemResponse.UpdatedAt); err != nil {
		return err
	}

	client.HealthProblems = append(client.HealthProblems, healthProblemResponse)

	return nil
}
