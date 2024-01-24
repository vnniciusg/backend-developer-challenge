package request

type UpdateHealthProblemRequestDTO struct {
	Grau int    `json:"grau" validate:"required,min=1,max=2" example:"1"`
	Name string `json:"name" validate:"required,min=3,max=255" example:"Dor de cabeça"`
}
