package clientrepository_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/clientrespository/clientrepositoryimpl"
	"github.com/vnniciusg/backend-developer-challenge/tests"
)

func TestCreateClient(t *testing.T) {

	db := tests.InitTestDB()

	clientRepository := clientrepositoryimpl.NewClientRepository(db)

	t.Run("should create a new client", func(t *testing.T) {

		parsedBirthDate, _ := date.ParseDate("01/01/2003")

		client := &entities.Client{
			Name:      "Vinnicius",
			BirthDate: parsedBirthDate,
			Sexo:      "m",
		}

		clientRepository.CreateClient(client)

		assert.NotNil(t, client)

	})

	t.Run("should return an error when the client name is empty", func(t *testing.T) {

		parsedBirthDate, _ := date.ParseDate("01/01/2003")

		client := &entities.Client{
			Name:      "",
			BirthDate: parsedBirthDate,
			Sexo:      "m",
		}

		err := clientRepository.CreateClient(client)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Nome do cliente não pode ser vazio")

	})

	t.Run("should return an error when the sex is not 'm' or 'f'", func(t *testing.T) {
		parsedBirthDate, _ := date.ParseDate("01/01/2000")

		client := &entities.Client{
			Name:      "Jane",
			BirthDate: parsedBirthDate,
			Sexo:      "x",
		}

		err := clientRepository.CreateClient(client)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Sexo do cliente deve ser 'm' ou 'f'")
	})
	t.Run("should return an error when the sex is empty", func(t *testing.T) {
		parsedBirthDate, _ := date.ParseDate("01/01/2000")

		client := &entities.Client{
			Name:      "Jane",
			BirthDate: parsedBirthDate,
			Sexo:      "",
		}

		err := clientRepository.CreateClient(client)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Sexo do cliente não pode ser vazio")
	})

	t.Run("should return an error when the birth date is in the future", func(t *testing.T) {
		parsedBirthDate, _ := date.ParseDate("01/01/2100")

		client := &entities.Client{
			Name:      "John",
			BirthDate: parsedBirthDate,
			Sexo:      "m",
		}

		err := clientRepository.CreateClient(client)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Data de nascimento do cliente não pode ser maior que a data atual")
	})
	t.Run("should return an error when the birth date is empty", func(t *testing.T) {

		client := &entities.Client{
			Name:      "John",
			BirthDate: time.Time{},
			Sexo:      "m",
		}

		err := clientRepository.CreateClient(client)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Data de nascimento do cliente não pode ser vazio")
	})

}
