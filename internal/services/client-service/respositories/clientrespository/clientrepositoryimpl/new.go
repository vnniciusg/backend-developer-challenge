package clientrepositoryimpl

import "database/sql"

type ClientRepositoryImpl struct {
	DB *sql.DB
}

func NewClientRepositoryImpl(db *sql.DB) *ClientRepositoryImpl {
	return &ClientRepositoryImpl{DB: db}
}
