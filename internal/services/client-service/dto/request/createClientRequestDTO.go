package request

type CreateClientRequestDTO struct {
	Name      string `json:"name" validate:"required,min=3,max=255" example:"John Doe"`
	BirthDate string `json:"birth_date" validate:"required" example:"01-01-1990"`
	Sexo      string `json:"sexo" validate:"required,oneof=f m" example:"m"`
}
