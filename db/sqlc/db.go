// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserByUUIDStmt, err = db.PrepareContext(ctx, getUserByUUID); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByUUID: %w", err)
	}
	if q.setUserVerifiedStmt, err = db.PrepareContext(ctx, setUserVerified); err != nil {
		return nil, fmt.Errorf("error preparing query SetUserVerified: %w", err)
	}
	if q.userEmailExistsStmt, err = db.PrepareContext(ctx, userEmailExists); err != nil {
		return nil, fmt.Errorf("error preparing query UserEmailExists: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserByUUIDStmt != nil {
		if cerr := q.getUserByUUIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByUUIDStmt: %w", cerr)
		}
	}
	if q.setUserVerifiedStmt != nil {
		if cerr := q.setUserVerifiedStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setUserVerifiedStmt: %w", cerr)
		}
	}
	if q.userEmailExistsStmt != nil {
		if cerr := q.userEmailExistsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userEmailExistsStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                  DBTX
	tx                  *sql.Tx
	createUserStmt      *sql.Stmt
	getUserByEmailStmt  *sql.Stmt
	getUserByUUIDStmt   *sql.Stmt
	setUserVerifiedStmt *sql.Stmt
	userEmailExistsStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                  tx,
		tx:                  tx,
		createUserStmt:      q.createUserStmt,
		getUserByEmailStmt:  q.getUserByEmailStmt,
		getUserByUUIDStmt:   q.getUserByUUIDStmt,
		setUserVerifiedStmt: q.setUserVerifiedStmt,
		userEmailExistsStmt: q.userEmailExistsStmt,
	}
}
