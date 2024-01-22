package clientcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http/responseshttp"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/validator"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
)

// @Summary Criar um novo cliente
// @Description Criar um novo cliente
// @Tags client
// @Accept json
// @Produce json
// @Param client body request.CreateClientRequestDTO true "Client"
// @Success 201 {object} responseshttp.RestSuccess
// @Failure 400 {object} responseshttp.RestErr
// @Failure 500 {object} responseshttp.RestErr
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

	err = cc.clientUseCase.CreateClient(request)

	if err != nil {
		restError := responseshttp.NewInternalServerError("Erro ao criar cliente")
		c.JSON(restError.Code, restError)
		return
	}

	restSuccess := responseshttp.NewCreated("Cliente criado com sucesso", nil)
	c.JSON(restSuccess.Code, restSuccess)
}
