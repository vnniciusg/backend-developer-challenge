package entities

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	Sexo      string    `json:"sexo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewClient(name string, birthDate time.Time, sexo string) *Client {

	return &Client{
		Name:      name,
		BirthDate: birthDate,
		Sexo:      sexo,
	}
}
