package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-pattern/internal/person"
	"go-pattern/internal/sqlops"
	"log"
)

func listTables(db *sql.DB) error {

	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tableName)
	}

	return nil
}

func main() {

	// Instantiate an instance of SQLite db handler
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Assign create SQLite Table function
	var sqliteTblCreator sqlops.TblCreatorFunc = person.CreateSQLiteTblFunc
	err = sqlops.CreateTable(context.TODO(), db, sqliteTblCreator)
	if err != nil {
		log.Fatal(err)
	}

	// Assign the name identifier method
	err = sqlops.CreateTable(context.TODO(), db, person.NameIdentifier{})
	if err != nil {
		log.Fatal(err)
	}

	err = listTables(db)
	if err != nil {
		log.Fatal(err)
	}
}
