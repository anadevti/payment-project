package main

import (
	"net/http"

	"github.com/anadevti/payment-project/internal"
	"github.com/anadevti/payment-project/internal/db"
	_ "github.com/lib/pq"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	r := internal.SetupRoutes(database)
	http.ListenAndServe(":3000", r)
}
