package persistence

import (
	"database/sql"
	"fmt"

	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/viper/configs"
)

func GetConnection() (*sql.DB, error) {
	configs, err := configs.LoadConfig()

	if err != nil {
		return nil, err
	}

	DBUrl := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		configs.DBUser, configs.DBPass, configs.DBHost, configs.DBPort, configs.DBName)

	db, err := sql.Open(configs.DBDriver, DBUrl)

	if err != nil {
		return nil, err
	}

	return db, nil
}
