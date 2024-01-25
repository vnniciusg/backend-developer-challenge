package tests

import (
	"database/sql"
	"fmt"

	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/viper/configs"
)

type SetupTestDatabase struct {
	DB *sql.DB
}

func NewSetupTestDatabase(db *sql.DB) *SetupTestDatabase {
	return &SetupTestDatabase{DB: db}
}

func (s *SetupTestDatabase) CreateTables() {
	s.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	s.DB.Exec(`CREATE TABLE IF NOT EXISTS tb_clients (
		id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY, name VARCHAR(255) NOT NULL, birth_date DATE NOT NULL, sexo CHAR(1) NOT NULL CHECK (sexo IN ('f', 'm', 'u')), created_at TIMESTAMPTZ DEFAULT current_timestamp, updated_at TIMESTAMPTZ DEFAULT current_timestamp
	);`)
	s.DB.Exec(`CREATE TABLE IF NOT EXISTS tb_health_problems (
		id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY, client_id uuid REFERENCES tb_clients (id), name VARCHAR(255) NOT NULL, grau INT CHECK (grau IN (1, 2)), created_at TIMESTAMPTZ DEFAULT current_timestamp, updated_at TIMESTAMPTZ DEFAULT current_timestamp
	);`)
}

func InitTestDB() *sql.DB {

	configs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	DBUrl := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		configs.TestDbUser, configs.TestDbPass, configs.TestDbHost, configs.TestDbPort, configs.TestDbName)

	testDB, err := sql.Open(configs.DBDriver, DBUrl)

	db := NewSetupTestDatabase(testDB)

	db.CreateTables()

	if err != nil {
		panic(err)
	}

	if err := testDB.Ping(); err != nil {
		panic(err)
	}

	return testDB
}
