package healthproblemusecase

import "github.com/vnniciusg/backend-developer-challenge/internal/services/client-service/respositories/healthproblemsrepository"

type HealthProblemUseCase struct {
	healthProblemRepository healthproblemsrepository.HealthProblemsRepository
}

func NewHealthProblemUseCase(healthProblemRepository healthproblemsrepository.HealthProblemsRepository) *HealthProblemUseCase {
	return &HealthProblemUseCase{
		healthProblemRepository: healthProblemRepository,
	}
}
