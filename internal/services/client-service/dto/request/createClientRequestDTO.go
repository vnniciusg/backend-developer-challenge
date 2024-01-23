package request

//swagget:model
type CreateClientRequestDTO struct {
	//Nome do cliente
	//required: true
	//min length: 3
	//max length: 255
	//example: John Doe
	Name string `json:"name" validate:"required,min=3,max=255"`

	//Data de nascimento do cliente (Formato : dd-mm-yyyy)
	//required: true
	//example: 01-01-2003
	BirthDate string `json:"birth_date" validate:"required"`

	//Sexo do cliente
	//required: true
	//enum: f, m
	//example: m
	Sexo string `json:"sexo" validate:"required,oneof=f m"`
}
