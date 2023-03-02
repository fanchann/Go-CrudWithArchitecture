package utils

import "database/sql"

func TransactionCommitAndRollBack(tx *sql.Tx) {
	if err := recover(); err != nil {
		err := tx.Rollback()
		ErrorWithPanic(err)
	} else {
		err := tx.Commit()
		ErrorWithPanic(err)
	}
}
