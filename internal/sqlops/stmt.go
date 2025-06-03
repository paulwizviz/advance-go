package sqlops

import (
	"context"
	"database/sql"
)

// StmtCreator is an interface to create SQL statement.
type StmtCreator interface {
	CreateStmt(db *sql.DB) (*sql.Stmt, []any, error)
}

// StmtCreatorFunc is a function that is a type of StmtWriter
// interface
type StmtCreatorFunc func(db *sql.DB) (*sql.Stmt, []any, error)

func (t StmtCreatorFunc) CreateStmt(db *sql.DB) (*sql.Stmt, []any, error) {
	return t(db)
}

// StmtMiddleware is a middleware function taking StmtCreators
type StmtMiddleware func(next StmtCreator) StmtCreator

// ChainStmtMiddleware is an function to create a chain of
// middlewares
func ChainStmtMiddle(handler StmtCreator, middlewares ...StmtMiddleware) StmtCreator {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// ExecInsert is a function to execute INSERT statement from
// a statement writer.
func ExecInsert(ctx context.Context, db *sql.DB, creator StmtCreator) error {
	stmt, args, err := creator.CreateStmt(db)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}
	return nil
}
