package request

import "time"

type CreateClientRequestDTO struct {
	Name           string                          `json:"name"`
	BirthDate      time.Time                       `json:"birth_date"`
	Sexo           string                          `json:"sexo"`
	HealthProblems []CreateHealthProblemRequestDTO `json:"health_problems"`
}
