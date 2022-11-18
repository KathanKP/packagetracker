package utilities

import (
	"fmt"

	"github.com/KathanKP/packagetracker/internal/domain"
)

func AddPackageToDB(packageToAdd domain.Package) error {
	fmt.Printf("Received request to add pacakge %+v.\n", packageToAdd)
	currentPackageStatus := GetLatestPackageStatus(packageToAdd.TrackingId, packageToAdd.Service)
	fmt.Println("Current package status that is fetched is: " + currentPackageStatus)

	// addStmt := `
	// INSERT INTO testtable (tacking_id, service, status, last_updated)
	// VALUES ($1, $2, $3, $4)`
	// err := database.ExecuteExecDBCommand(addStmt, packageToAdd.TrackingId, packageToAdd.Service, currentPackageStatus, packageToAdd)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("Added the package to DB: %+v.\n", packageToAdd)
	return nil
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
