package sqlops

import (
	"context"
	"database/sql"
	"fmt"
)

type data struct {
	Value string
}

var (
	initDBFunc StmtWriterFunc = func(db *sql.DB) (*sql.Stmt, []any, error) {
		stmt, err := db.Prepare(`INSERT INTO data (value) VALUES( ? )`)
		if err != nil {
			return nil, nil, err
		}
		d := data{
			Value: "abc",
		}
		return stmt, []any{d.Value}, nil
	}
)

func Example_writeTblStmtFunc() {
	db, _ := NewSQLiteMem()
	defer db.Close()
	db.Exec(`CREATE TABLE data(
	   pk INTEGER PRIMARY KEY,
	   value TEXT)`)

	err := ExecInsert(context.TODO(), db, initDBFunc)
	if err != nil {
		fmt.Println(err)
	}

	result, err := db.Query("SELECT value FROM data")
	if err != nil {
		fmt.Println(err)
	}

	for result.Next() {
		var value string
		result.Scan(&value)
		fmt.Println(value)
	}

	// Output:
	// abc
}
