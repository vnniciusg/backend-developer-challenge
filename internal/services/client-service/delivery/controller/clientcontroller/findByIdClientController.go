package clientcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http/responseshttp"
)

// FindClientById godoc
// @Summary Busca um cliente pelo id
// @Description Busca um cliente pelo id
// @Tags Clientes
// @Accept  json
// @Produce  json
// @Param id path string true "Id do cliente"
// @Success 200 {object} responseshttp.RestSuccess
// @Failure 400 {object}  responseshttp.RestErr
// @Failure 404 {object}  responseshttp.RestErr
// @Failure 500 {object}  responseshttp.RestErr
// @Router /api/v1/clients/{id} [get]
func (cc *ClientController) FindClientById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		restErr := responseshttp.NewInternalServerError("Falha ao converter id")
		c.JSON(restErr.Code, restErr)
		return
	}

	if id == uuid.Nil {
		restErr := responseshttp.NewBadRequestErr("Id inválido")
		c.JSON(restErr.Code, restErr)
		return
	}

	client, err := cc.clientUseCase.FindClientById(id)

	if err != nil {
		restErr := responseshttp.NewInternalServerError("Falha ao buscar cliente")
		c.JSON(restErr.Code, restErr)
		return
	}

	if client == nil {
		restErr := responseshttp.NewNotFoundError("Cliente não encontrado")
		c.JSON(restErr.Code, restErr)
		return
	}

	restSuccess := responseshttp.NewOk("Cliente encontrado", client)
	c.JSON(restSuccess.Code, restSuccess)
}
