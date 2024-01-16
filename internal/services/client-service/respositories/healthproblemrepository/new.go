package healthproblemrepository

import "database/sql"

type HealthProblemsRepository struct {
	DB *sql.DB
}

func NewHealthProblemsRepository(db *sql.DB) *HealthProblemsRepository {
	return &HealthProblemsRepository{DB: db}
}
