package server

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/KathanKP/packagetracker/internal/domain"
	"github.com/KathanKP/packagetracker/internal/services/fedex"
)

func addPackageHandler(w http.ResponseWriter, r *http.Request) {
	packageToAdd := domain.Package{}
	err := json.NewDecoder(r.Body).Decode(&packageToAdd)
	if err != nil {
		panic(err)
	}
	packageToAdd.LastUpdated = time.Now().Local()
	// call API To fetch current status
	fedex.GetPackageStatus("id")
	packageToAdd.CurrentStatus = "IN PROGRESS"
	fmt.Printf("Sending request to DB to add pacakge %+v.\n", packageToAdd)

	err = AddPackage(packageToAdd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Processed and added the package successfully!")
}

func main() {
	// Check and create a table if not already created
	CreateTable()

	// Register the addPackageHandler function
	http.HandleFunc("/addpackage", addPackageHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":"+dbPort, nil); err != nil {
		log.Fatal(err)
	}
}
