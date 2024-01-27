package create_client_repository_test

import (
	"testing"

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

}
