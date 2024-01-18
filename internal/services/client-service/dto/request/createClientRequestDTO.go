package request

type CreateClientRequestDTO struct {
	Name           string                           `json:"name"`
	BirthDate      string                           `json:"birth_date"`
	Sexo           string                           `json:"sexo"`
	HealthProblems []*CreateHealthProblemRequestDTO `json:"health_problems"`
}
