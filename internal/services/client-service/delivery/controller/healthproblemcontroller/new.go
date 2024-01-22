package healthproblemcontroller

import "github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/usecase/healthproblemusecase"

type HealProblemsController struct {
	healthProblemUseCase *healthproblemusecase.HealthProblemUseCase
}

func NewHealProblemsController(healthProblemUseCase *healthproblemusecase.HealthProblemUseCase) *HealProblemsController {
	return &HealProblemsController{
		healthProblemUseCase: healthProblemUseCase,
	}
}
