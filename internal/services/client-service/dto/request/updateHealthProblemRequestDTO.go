package request

type UpdateHealthProblemRequestDTO struct {
	Grau int    `json:"grau" validate:"min=1,max=2" example:"1"`
	Name string `json:"name" validate:"min=3,max=255" example:"Dor de cabeça"`
}
