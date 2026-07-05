package main

import (
	"net/http"

	"github.com/anadevti/payment-project/internal"
)

func main() {
	r := internal.SetupRoutes()
	http.ListenAndServe(":3000", r)
}
