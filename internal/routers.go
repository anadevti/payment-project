package internal

import (
	"github.com/anadevti/payment-project/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/accounts/{id}", handlers.GetAccount)
		r.Get("/accounts/", handlers.SomeDefaultHandler)
	})
	return r
}
