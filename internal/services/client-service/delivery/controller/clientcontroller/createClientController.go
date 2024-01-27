package clientcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http/responseshttp"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/validator"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

// @Summary Criar um novo cliente
// @Description Cria um novo cliente
// @Tags Clientes
// @Accept json
// @Produce json
// @Param client body request.CreateClientRequestDTO true "Estrutura de dados para criar um novo cliente"
// @Success 201 {object} responseshttp.RestSuccess "Cliente criado com sucesso"
// @Failure 400 {object} responseshttp.RestErr "Erro ao processar a solicitação"
// @Failure 500 {object} responseshttp.RestErr "Erro interno do servidor"
// @Router /api/v1/clients [post]
func (cc *ClientController) CreateClient(c *gin.Context) {

	var request *request.CreateClientRequestDTO

	err := c.ShouldBindJSON(&request)

	if err != nil {
		restError := responseshttp.NewInternalServerError("Erro ao realizar conversão do JSON")
		c.JSON(restError.Code, restError)
		return
	}

	validation := validator.ValidateDataRequest(request)

	if validation != nil {
		restError := responseshttp.NewBadRequestValidationError("Erro ao validar dados", validation)
		c.JSON(restError.Code, restError)
		return
	}

	birthDate, err := date.ParseDate(request.BirthDate)

	if err != nil {
		restError := responseshttp.NewInternalServerError(err.Error())
		c.JSON(restError.Code, restError)
		return
	}

	newClient := entities.NewClient(request.Name, birthDate, request.Sexo)

	if err != nil {
		restError := responseshttp.NewBadRequestErr(err.Error())
		c.JSON(restError.Code, restError)
		return
	}

	err = cc.clientUseCase.CreateClient(newClient)

	if err != nil {
		restError := responseshttp.NewInternalServerError("Erro ao criar cliente")
		c.JSON(restError.Code, restError)
		return
	}

	restSuccess := responseshttp.NewCreated("Cliente criado com sucesso", nil)
	c.JSON(restSuccess.Code, restSuccess)
}
