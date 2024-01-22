package di

import (
	"database/sql"

	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/delivery/controller/healthproblemcontroller"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/healthproblemsrepository/healthproblemrepositoryimpl"
	"github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/usecase/healthproblemusecase"
)

func ConfigHealthProblemDI(conn *sql.DB) *healthproblemcontroller.HealProblemsController {
	healthProblemRepository := healthproblemrepositoryimpl.NewHealthProblemRepository(conn)
	healthProblemUseCase := healthproblemusecase.NewHealthProblemUseCase(healthProblemRepository)
	healthProblemController := healthproblemcontroller.NewHealProblemsController(healthProblemUseCase)

	return healthProblemController

}
