package person

import (
	"fmt"
	"go-pattern/internal/sqlops"
	"strings"
)

// NameIdentifier is a data type representing a named based
// identifier with struct tages representing fields in
// a table `name_id`,
type NameIdentifier struct {
	PK        int    `json:"pk" sqlite:"pk,INTEGER,PRIMARY_KEY"`
	FirstName string `json:"first_name" sqlite:"firstname,TEXT"`
	Surname   string `json:"surname" sqlite:"surname,TEXT"`
}

// CreateStmt is a concrete implementation of sqlops.TblCreator
// interface
func (n NameIdentifier) CreateStmt() string {
	tags := sqlops.ExtractTags("sqlite", n)
	createStmt := "CREATE TABLE IF NOT EXISTS name_id ("
	for _, t := range tags {
		if strings.Contains(t.Tag, "pk") {
			tag := strings.ReplaceAll(t.Tag, ",", " ")
			tag = strings.Replace(tag, "_", " ", -1)
			createStmt = fmt.Sprintf("%s %s,", createStmt, tag)
			continue
		}
		createStmt = fmt.Sprintf("%s %s,", createStmt, strings.ReplaceAll(t.Tag, ",", " "))
	}
	createStmt = createStmt[:len(createStmt)-1]
	createStmt = fmt.Sprintf("%s)", createStmt)
	return createStmt
}

// CreateSQLiteTblFunc is a function type that is
// an implementation of sqlops.TblCreatorFunc, which
// is also a type of sqlops.TblCreator
func CreateSQLiteTblFunc() string {
	return `CREATE TABLE IF NOT EXISTS assigned_id (
	   pk INTEGER PRIMARY KEY, 
	   email TEXT);
	   
	CREATE TABLE IF NOT EXISTS person(
	   pk INTEGER PRIMARY KEY,
	   a_id INTEGER,
       FOREIGN KEY(a_id) REFERENCES assigned_id(pk)
	);`
}
