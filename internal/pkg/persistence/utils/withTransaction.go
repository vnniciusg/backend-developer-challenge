package utils

import "database/sql"

func WithTransaction(db *sql.DB, transactionFunc func(*sql.Tx) error) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	return transactionFunc(tx)
}
