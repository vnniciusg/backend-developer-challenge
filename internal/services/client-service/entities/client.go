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
