package sqlops

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

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

type StructTag struct {
	FieldName string
	Tag       string
}

// ExtractTags extract struct ExtractTags of direct fields
// it will not extract ExtractTags from composed
// fields
func ExtractTags(tagName string, typ any) []StructTag {
	ev := reflect.Indirect(reflect.ValueOf(typ))
	tags := []StructTag{}
	for i := 0; i < ev.Type().NumField(); i++ {
		tag := StructTag{}
		tag.FieldName = ev.Type().Field(i).Name
		t := ev.Type().Field(i).Tag
		tElems := strings.Split(string(t), " ")
		for _, tElem := range tElems {
			if strings.Contains(tElem, tagName) {
				sElems := strings.Split(tElem, ":")
				tag.Tag = sElems[1][1 : len(sElems[1])-1]
			}
		}
		tags = append(tags, tag)
	}
	return tags
}
