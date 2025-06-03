package sqlops

import (
	"context"
	"database/sql"
	"fmt"
)

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
