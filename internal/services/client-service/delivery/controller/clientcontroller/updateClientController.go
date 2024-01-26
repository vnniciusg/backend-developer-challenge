package clientcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/date"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http/responseshttp"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

// @Summary Atualizar um cliente
// @Description Atualizar um cliente pelo id
// @Tags Clientes
// @Accept  json
// @Produce  json
// @Param id path string true "Client ID"
// @Param cliente body request.UpdateClientRequestDTO true "Cliente"
// @Success 200 {object} responseshttp.RestSuccess "Cliente atualizado com sucesso"
// @Failure 400 {object} responseshttp.RestErr "Erro de validação"
// @Failure 500 {object} responseshttp.RestErr "Falha ao criar cliente"
// @Router /api/v1/clients/{id} [patch]
func (cc *ClientController) UpdateClient(c *gin.Context) {

	var request *request.UpdateClientRequestDTO

	err := c.ShouldBindJSON(&request)

	if err != nil {
		restErr := responseshttp.NewInternalServerError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	id, err := uuid.Parse(c.Param("id"))

	if id == uuid.Nil {
		restErr := responseshttp.NewBadRequestErr("Id inválido")
		c.JSON(restErr.Code, restErr)
		return
	}

	if err != nil {
		restErr := responseshttp.NewBadRequestErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	parsedBirthDate, err := date.ParseDate(request.BirthDate)

	if err != nil {
		restErr := responseshttp.NewInternalServerError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	client := entities.NewClient(request.Name, parsedBirthDate, request.Sexo)

	err = cc.clientUseCase.UpdateClient(client)

	if err != nil {
		restErr := responseshttp.NewInternalServerError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	restSuccess := responseshttp.NewOk("Cliente atualizado com sucesso", nil)
	c.JSON(restSuccess.Code, restSuccess)

}
