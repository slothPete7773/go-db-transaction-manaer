package typesafetransactionproviderwithgeneric

import (
	"context"
	"database/sql"
	"fmt"
)

type impl[T any] struct {
	db *sql.DB
	wt withTx[T]
}

func NewTrm[T withTx[T]](db *sql.DB, wt T) *impl[T] {
	return &impl[T]{
		db: db,
		wt: wt,
	}
}

func (t *impl[T]) InTx(ctx context.Context, fn func(repo T) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	repo := t.wt.WithTx(tx)
	if err := fn(repo); err != nil {
		return fmt.Errorf("tx callback: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
