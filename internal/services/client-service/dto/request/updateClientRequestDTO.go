package request

type UpdateClientRequestDTO struct {
	Name      string `json:"name" validate:"min=3,max=255" example:"John Doe"`
	BirthDate string `json:"birth_date" example:"01-01-1990"`
	Sexo      string `json:"sexo" validate:"oneof=f m" example:"m"`
}
