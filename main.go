package main

import (
	"net/http"

	"github.com/anadevti/payment-project/internal"
	"github.com/anadevti/payment-project/internal/db"
)

func main() {
	r := internal.SetupRoutes()
	http.ListenAndServe(":3000", r)

	database := db.ConnectDB()
	defer database.Close()
}
