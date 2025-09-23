package main

import (
	"database/sql"
	"log"
)

type DBTransactor struct {
	tx *sql.Tx
}

func NewTransactor(db *sql.DB) DBTransactor {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	return DBTransactor{tx: tx}
}

func (t *DBTransactor) Begin() *Tx {
	return &Tx{tx: t.tx, commit: false}
}

type Tx struct {
	tx     *sql.Tx
	commit bool
}

func (t *Tx) Instance() *sql.Tx {
	return t.tx
}

func (t *Tx) Commit() error {
	t.commit = true
	return t.tx.Commit()
}

func (t *Tx) SafeRollback(recoverErr any) {
	if !t.commit {
		t.tx.Rollback()
		return
	}
	if recoverErr != nil {
		t.tx.Rollback()
		return
	}
}
