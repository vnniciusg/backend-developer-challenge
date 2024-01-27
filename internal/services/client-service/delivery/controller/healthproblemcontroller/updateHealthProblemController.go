package healthproblemcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http/responseshttp"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

// @Summary Atualizar um problema de saúde
// @Description Atualizar um problema de saúde pelo id
// @Tags Health Problem
// @Accept  json
// @Produce  json
// @Param id path string true "Health Problem ID"
// @Param healthProblem body request.UpdateHealthProblemRequestDTO true "Health Problem"
// @Success 200 {object} responseshttp.RestSuccess "Problema de saúde atualizado com sucesso"
// @Failure 400 {object} responseshttp.RestErr "Erro de validação"
// @Failure 500 {object} responseshttp.RestErr "Falha ao criar problema de saúde"
// @Router /api/v1/health-problems/{id} [patch]
func (hpc *HealProblemsController) UpdateHealthProblem(c *gin.Context) {
	var request *request.UpdateHealthProblemRequestDTO

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
		restErr := responseshttp.NewInternalServerError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	healthProblem := entities.NewHealthProblem(request.Name, id, request.Grau)

	err = hpc.healthProblemUseCase.UpdateHealthProblem(healthProblem)

	if err != nil {
		restErr := responseshttp.NewInternalServerError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	restSuccess := responseshttp.NewOk("Problema de saúde atualizado com sucesso", nil)
	c.JSON(restSuccess.Code, restSuccess)

}
