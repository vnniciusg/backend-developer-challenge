package repositories

import (
	"database/sql"
)

type SetupTestDatabase struct {
	DB *sql.DB
}

func (s *SetupTestDatabase) Setup() {
	s.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	s.DB.Exec(`CREATE TABLE IF NOT EXISTS tb_clients (
		id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY, name VARCHAR(255) NOT NULL, birth_date DATE NOT NULL, sexo CHAR(1) NOT NULL CHECK (sexo IN ('f', 'm', 'u')), created_at TIMESTAMPTZ DEFAULT current_timestamp, updated_at TIMESTAMPTZ DEFAULT current_timestamp
	);`)
	s.DB.Exec(`CREATE TABLE IF NOT EXISTS tb_health_problems (
		id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY, client_id uuid REFERENCES tb_clients (id), name VARCHAR(255) NOT NULL, grau INT CHECK (grau IN (1, 2)), created_at TIMESTAMPTZ DEFAULT current_timestamp, updated_at TIMESTAMPTZ DEFAULT current_timestamp
	);`)
}

func (s *SetupTestDatabase) TearDown() {
	s.DB.Exec(`DROP TABLE IF EXISTS tb_health_problems;`)
	s.DB.Exec(`DROP TABLE IF EXISTS tb_clients;`)

	s.DB.Close()
}

func NewSetupTestDatabase(db *sql.DB) *SetupTestDatabase {
	return &SetupTestDatabase{DB: db}
}

func setupTestDB() *sql.DB {
	testDB, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/clients?sslmode=disable")
	if err != nil {
		panic(err)
	}

	setupDB := NewSetupTestDatabase(testDB)

	setupDB.Setup()

	return testDB
}

func tearDownTestDB(testDB *sql.DB) {
	setupDB := NewSetupTestDatabase(testDB)
	setupDB.TearDown()
}
