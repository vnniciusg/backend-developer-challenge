package clientcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http/responseshttp"
)

// FindAllClients godoc
// @Summary Busca todos os clientes
// @Description Busca todos os clientes
// @Tags Clientes
// @Accept  json
// @Produce  json
// @Success 200 {object} responseshttp.RestSuccess
// @Failure 400 {object} responseshttp.RestErr
// @Failure 500 {object} responseshttp.RestErr
// @Router /api/v1/clients [get]
func (cc *ClientController) FindAllClients(c *gin.Context) {

	clients, err := cc.clientUseCase.FindAllClients()

	if err != nil {
		restErr := responseshttp.NewInternalServerError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	if clients == nil {
		restErr := responseshttp.NewNotFoundError("Nenhum cliente encontrado")
		c.JSON(restErr.Code, restErr)
		return
	}

	restSuccess := responseshttp.NewOk("Clientes encontrados com sucesso", clients)
	c.JSON(restSuccess.Code, restSuccess)
}
