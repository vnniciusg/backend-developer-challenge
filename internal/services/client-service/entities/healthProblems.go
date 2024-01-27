package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type HealthProblem struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ClientId  uuid.UUID `json:"client_id"`
	Grau      int       `json:"grau"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewHealthProblem(name string, clientId uuid.UUID, grau int) *HealthProblem {
	return &HealthProblem{
		Name:     name,
		ClientId: clientId,
		Grau:     grau,
	}
}

func (hp *HealthProblem) ValidateHealthProblem() error {
	if hp.Name == "" {
		return errors.New("Nome do problema de saúde não pode ser vazio")
	}

	if hp.ClientId == uuid.Nil {
		return errors.New("Id do cliente não pode ser vazio")
	}

	if hp.Grau != 1 && hp.Grau != 2 {
		if hp.Grau == 0 {
			return errors.New("Grau do problema de saúde não pode ser vazio")
		}
		return errors.New("Grau do problema de saúde deve ser 1 ou 2")
	}

	return nil
}
