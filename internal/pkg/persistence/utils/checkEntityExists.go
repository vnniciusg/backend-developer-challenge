package utils

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type EntityType string

const (
	HealthProblem EntityType = "HealthProblem"
	Client        EntityType = "Client"
)

func CheckEntityExists(db *sql.DB, query string, id uuid.UUID, entityType EntityType) error {
	_, err := db.Exec(query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New(string(entityType) + " com id " + id.String() + " não existe")
		}
		return err
	}

	return nil

}
