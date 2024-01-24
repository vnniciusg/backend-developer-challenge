package clientrepository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/clientrespository/clientrepositoryimpl"
	"github.com/vnniciusg/backend-developer-challenge/tests/repositories"
)

func TestCreateClient(t *testing.T) {

	db := repositories.SetupTestDB()

	defer repositories.TearDownTestDB(db)

	clientRepository := clientrepositoryimpl.NewClientRepository(db)

	t.Run("should create a new client", func(t *testing.T) {

		client := &request.CreateClientRequestDTO{
			Name:      "Vinicius",
			BirthDate: "01-01-2003",
			Sexo:      "m",
		}

		err := clientRepository.CreateClient(client)

		if err != nil {
			assert.Nil(t, err)
		}

		assert.NotNil(t, client)

	})

	t.Run("should return an error when the client is invalid", func(t *testing.T) {

		client := &request.CreateClientRequestDTO{
			Name:      "",
			BirthDate: "01-01-2003",
			Sexo:      "m",
		}

		err := clientRepository.CreateClient(client)

		assert.NotNil(t, err)

	})

}
