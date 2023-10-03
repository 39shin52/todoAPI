package transaction

import (
	"context"
	"database/sql"
	"log"
)

type Transaction struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) *Transaction {
	return &Transaction{db: db}
}

func (t *Transaction) Transaction(ctx context.Context, f func(ctx context.Context, tx *sql.Tx) (interface{}, error)) (interface{}, error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("failed to MySQL Rollback: %v", err)
			}

			panic(p)
		}

		if err != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("failed to MySQL Rollback: %v", err)
			}

			return
		}

		if err := tx.Commit(); err != nil {
			log.Printf("failed to MySQL Commit: %v", err)
		}
	}()

	data, err := f(ctx, tx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
