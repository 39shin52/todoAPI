package transaction

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type TxAdmin struct {
	db *sql.DB
}

func NewTxRepository(db *sql.DB) *TxAdmin {
	return &TxAdmin{db: db}
}

func (t *TxAdmin) Transaction(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
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

	if err := f(ctx); err != nil {
		return fmt.Errorf("transaction query failed: %w", err)
	}

	return nil
}
