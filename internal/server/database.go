package server

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/KathanKP/packagetracker/internal/domain"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "testdb"
	password = "db1234"
	dbname   = "testdb"
	dbPort   = "8080"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func CreateTable() {
	fmt.Println("Table not found. Creating a new one...")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	createStmt := `CREATE TABLE IF NOT EXISTS testtable (
		tacking_id VARCHAR (255) PRIMARY KEY,
		service VARCHAR (255) NOT NULL,
		status VARCHAR (255) NOT NULL,
		last_updated TIMESTAMP NOT NULL)`
	_, err = db.Exec(createStmt)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully created and connected to table")
}

// func FetchPackageFromDBHandler(w http.ResponseWriter, r *http.Request) {
// 	queryParams := make(map[string]string)
// 	for k, v := range r.URL.Query() {
// 		queryParams[k] = v[0]
// 	}
// 	fmt.Printf("Received request to fetch package with id %s.\n", queryParams["id"])

// 	var trackingId string
// 	var service string
// 	var status string
// 	var timeup string
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	selectStmt := `SELECT * from testtable`
// 	row := db.QueryRow(selectStmt)
// 	err = row.Scan(&trackingId, &service, &status, &timeup)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Fetched rows from table")
// 	fmt.Println(trackingId, service, status, timeup)
// }

func AddPackage(packageToAdd domain.Package) error {
	fmt.Printf("Received DB request to add pacakge %+v.\n", packageToAdd)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()
	addStmt := `
	INSERT INTO testtable (tacking_id, service, status, last_updated)
	VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(addStmt, packageToAdd.TrackingId,
		packageToAdd.Service,
		packageToAdd.CurrentStatus,
		packageToAdd.LastUpdated)
	if err != nil {
		return err
	}
	fmt.Printf("Added the package to DB: %+v.\n", packageToAdd)
	return nil
}
