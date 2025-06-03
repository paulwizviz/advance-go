package sqlops

import (
	"context"
	"database/sql"
)

// StmtWriter is an interface to write data to
// an SQL database
type StmtWriter interface {
	WriteStmt(db *sql.DB) (*sql.Stmt, []any, error)
}

// StmtWriterFunc is a function that is a type of StmtWriter
// interface
type StmtWriterFunc func(db *sql.DB) (*sql.Stmt, []any, error)

func (t StmtWriterFunc) WriteStmt(db *sql.DB) (*sql.Stmt, []any, error) {
	return t(db)
}

// ExecInsert is a function to execute INSERT statement from
// a statement writer.
func ExecInsert(ctx context.Context, db *sql.DB, writer StmtWriter) error {
	stmt, args, err := writer.WriteStmt(db)
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
