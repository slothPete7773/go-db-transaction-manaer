package txdemo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	// "github.com/pkg/errors"
	// "go.uber.org/multierr"
)

type DB interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type TxManager interface {
	Run(
		ctx context.Context,
		callback func(ctx context.Context, tx *sql.Tx) error,
	) error
}

type TxKey string

var CtxWithTx = TxKey("tx")

type SQLTransactionManager struct {
	db *sql.DB
}

func NewTxManager(db *sql.DB) *SQLTransactionManager {
	return &SQLTransactionManager{db: db}
}

func (m *SQLTransactionManager) Run(ctx context.Context, callback func(ctx context.Context, tx *sql.Tx) error) (rErr error) {
	tx, err := m.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	log.Println("begin tx")

	defer func() {
		if rErr != nil {
			tx.Rollback()
			rErr = err
			// rErr = multierr.Combine(rErr, errors.WithStack(tx.Rollback()))
		}
	}()

	defer func() {
		if rec := recover(); rec != nil {
			if e, ok := rec.(error); ok {
				rErr = e
			} else {
				err = fmt.Errorf("%s", rec)
			}
		}
	}()

	log.Println("do calback")
	if err = callback(ctx, tx); err != nil {
		log.Println("callback err:", err.Error())
		return err
	}

	log.Println("commit")
	return tx.Commit()
}
