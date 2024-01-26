package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type HealthProblems struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ClientId  uuid.UUID `json:"client_id"`
	Grau      int       `json:"grau"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewHealthProblems(name string, clientId uuid.UUID, grau int) *HealthProblems {
	return &HealthProblems{
		Name:     name,
		ClientId: clientId,
		Grau:     grau,
	}
}

func ValidateHealthProblems(healthProblems *HealthProblems) error {
	if healthProblems.Name == "" {
		return errors.New("Nome do problema de saúde não pode ser vazio")
	}

	if healthProblems.ClientId == uuid.Nil {
		return errors.New("Id do cliente não pode ser vazio")
	}

	if healthProblems.Grau != 1 && healthProblems.Grau != 2 {
		if healthProblems.Grau == 0 {
			return errors.New("Grau do problema de saúde não pode ser vazio")
		}
		return errors.New("Grau do problema de saúde deve ser 1 ou 2")
	}

	return nil
}
