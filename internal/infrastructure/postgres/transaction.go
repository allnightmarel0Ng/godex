package postgres

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type Transaction interface {
	Commit() error
	Rollback() error
	QueryRow(sql string, args ...interface{}) pgx.Row
	Exec(sql string, args ...interface{}) (pgconn.CommandTag, error)
}

type transaction struct {
	tx pgx.Tx
	ctx context.Context
}

func newTransaction(ctx context.Context, tx pgx.Tx) Transaction {
	return &transaction{
		tx: tx,
		ctx: ctx,
	}
}

func (t *transaction) Commit() error {
	return t.tx.Commit(t.ctx)
}

func (t *transaction) Rollback() error {
	return t.tx.Rollback(t.ctx)
}

func (t *transaction) QueryRow(sql string, args ...interface{}) pgx.Row {
	return t.tx.QueryRow(t.ctx, sql, args...)
}

func (t *transaction) Exec(sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return t.tx.Exec(t.ctx, sql, args...)
}
