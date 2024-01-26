package entities

import (
	"errors"
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

func ValidateClient(client *Client) error {
	if client.Name == "" {
		return errors.New("Nome do cliente não pode ser vazio")
	}

	if client.Sexo == "" {
		return errors.New("Sexo do cliente não pode ser vazio")
	}

	if client.BirthDate.IsZero() {
		return errors.New("Data de nascimento do cliente não pode ser vazio")
	}

	return nil
}
