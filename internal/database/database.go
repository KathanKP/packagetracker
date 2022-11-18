package database

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "testdb"
	password  = "db1234"
	dbname    = "testdb"
	dbPort    = "8080"
	tableName = "testtable"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func getSqlObject() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func CreateTable() {
	fmt.Println("Table not found. Creating a new one...")
	createStmt := `CREATE TABLE IF NOT EXISTS testtable (
		tacking_id VARCHAR (255) PRIMARY KEY,
		service VARCHAR (255) NOT NULL,
		status VARCHAR (255) NOT NULL,
		last_updated TIMESTAMP NOT NULL)`
	err := ExecuteExecDBCommand(createStmt)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully created table")
}

func ExecuteExecDBCommand(dbExecCmd string, cmdArgs ...any) error {
	fmt.Println("Executing the command: " + dbExecCmd)
	db := getSqlObject()
	_, err := db.Exec(dbExecCmd, cmdArgs)
	db.Close()
	return err
}
