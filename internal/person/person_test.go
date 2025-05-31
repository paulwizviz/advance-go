package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStmtMethod(t *testing.T) {
	expected := "CREATE TABLE IF NOT EXISTS name_id ( pk INTEGER PRIMARY KEY, firstname TEXT, surname TEXT)"
	pi := NameIdentifier{}
	actual := pi.CreateStmt()
	assert.Equal(t, expected, actual)
}

func TestCreateSQLiteTblStmt(t *testing.T) {
	expected := `CREATE TABLE IF NOT EXISTS assigned_id (
	   pk INTEGER PRIMARY KEY, 
	   email TEXT);
	   
	CREATE TABLE IF NOT EXISTS person(
	   pk INTEGER PRIMARY KEY,
	   a_id INTEGER,
       FOREIGN KEY(a_id) REFERENCES assigned_id(pk)
	);`
	actual := CreateSQLiteTblFunc()
	assert.Equal(t, expected, actual)
}
