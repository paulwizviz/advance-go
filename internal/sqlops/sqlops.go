package sqlops

import (
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

// NewSQLiteMem create an instance of a
// memory based SQLite DB handler.
func NewSQLiteMem() (*sql.DB, error) {
	db, err := sql.Open(sqliteVer, ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w:%v", ErrConn, err)
	}
	return db, nil
}

// NewSQLiteFile create an instance of a file base
// SQLite DB handler.
func NewSQLiteFile(f string) (*sql.DB, error) {
	db, err := sql.Open(sqliteVer, f)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}

// NewPostgres creates an instance of a Postgres DB handlers
func NewPostgres(username string, password string, host string, port uint, dbname string) (*sql.DB, error) {
	connStmt := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", connStmt)
	if err != nil {
		return nil, err
	}
	return db, nil
}
