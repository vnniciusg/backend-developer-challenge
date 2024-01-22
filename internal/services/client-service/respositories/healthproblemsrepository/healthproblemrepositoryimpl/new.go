package healthproblemrepositoryimpl

import "database/sql"

type HealthProblemRepository struct {
	DB *sql.DB
}

func NewHealthProblemRepository(db *sql.DB) *HealthProblemRepository {
	return &HealthProblemRepository{
		DB: db,
	}
}
