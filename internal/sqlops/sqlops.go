package sqlops

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrConn  = errors.New("db connect error")
	ErrStmt  = errors.New("statement error")
	ErrTable = errors.New("create table error")
)

const (
	sqliteVer = "sqlite3"
)

func NewSQLiteMem() (*sql.DB, error) {
	db, err := sql.Open(sqliteVer, ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w:%v", ErrConn, err)
	}
	return db, nil
}

// TblCreator is a ainglw interface representing
// operations for creating table.
type TblCreator interface {
	CreateStmt() string
}

// TblCreatorFunc is a function that is a type of TblCreator
type TblCreatorFunc func() string

func (t TblCreatorFunc) CreateStmt() string {
	return t()
}

// CreateTable is a wrapper function to create an SQL table based on
// TblCreator interface as arguument supplied.
func CreateTable(ctx context.Context, db *sql.DB, creator TblCreator) error {
	stmt := creator.CreateStmt()
	_, err := db.ExecContext(ctx, stmt)
	if err != nil {
		return fmt.Errorf("%w:%v", ErrTable, err)
	}
	return nil
}

// TblWriter is an interface to write data to
// an SQL database
type TblWriter interface {
	Write(db *sql.DB) (*sql.Stmt, []any, error)
}

func WriteTable(ctx context.Context, db *sql.DB, writer TblWriter) error {
	stmt, args, err := writer.Write(db)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args)
	if err != nil {
		return err
	}
	return nil
}
