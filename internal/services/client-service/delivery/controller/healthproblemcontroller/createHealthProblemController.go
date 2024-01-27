package healthproblemcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http/responseshttp"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/validator"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/dto/request"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/entities"
)

// @Summary Criar um prbolema de saúde
// @Description Criar um prbolema de saúde
// @Tags Health Problem
// @Accept  json
// @Produce  json
// @Param id path string true "Client id"
// @Param HealthProblem body []request.CreateHealthProblemRequestDTO true "HealthProblem body"
// @Success 201 {object} responseshttp.RestSuccess "Problema(s) de saúde criado(s) com sucesso"
// @Failure 400 {object} responseshttp.RestErr "Erro de validação"
// @Failure 500 {object} responseshttp.RestErr "Falha ao criar problema de saúde"
// @Router /api/v1/health-problems/{clientId} [post]
func (hpc *HealProblemsController) CreateHealthProblem(c *gin.Context) {

	clientId, err := uuid.Parse(c.Param("clientId"))

	if err != nil {
		restErr := responseshttp.NewInternalServerError("Falha ao converter id")
		c.JSON(restErr.Code, restErr)
		return
	}

	if clientId == uuid.Nil {
		restErr := responseshttp.NewBadRequestErr("Id inválido")
		c.JSON(restErr.Code, restErr)
		return
	}

	var request []request.CreateHealthProblemRequestDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := responseshttp.NewBadRequestErr("Falha ao converter body")
		c.JSON(restErr.Code, restErr)
		return
	}

	var healthProblems []*entities.HealthProblem
	for _, healthProblem := range request {
		newHealthProblem, err := validateAndConvertHealthProblemEntityRequest(clientId, &healthProblem)
		if err != nil {
			c.JSON(err.Code, err)
			return
		}
		healthProblems = append(healthProblems, newHealthProblem)
	}

	err = hpc.healthProblemUseCase.CreateHealthProblem(healthProblems)

	if err != nil {
		restErr := responseshttp.NewInternalServerError("Falha ao criar problema de saúde")
		c.JSON(restErr.Code, restErr)
		return
	}

	restSuccess := responseshttp.NewCreated("Problema(s) de saúde criado(s) com sucesso", nil)
	c.JSON(restSuccess.Code, restSuccess)

}

func validateAndConvertHealthProblemEntityRequest(id uuid.UUID, request *request.CreateHealthProblemRequestDTO) (*entities.HealthProblem, *responseshttp.RestErr) {

	validator := validator.ValidateDataRequest(request)
	if validator != nil {
		return nil, responseshttp.NewBadRequestValidationError("Erro de validação", validator)
	}

	healthProblem := entities.NewHealthProblem(request.Name, id, request.Grau)

	return healthProblem, nil
}
