package internal

import (
	"database/sql"

	"github.com/anadevti/payment-project/internal/handlers"
	"github.com/anadevti/payment-project/internal/repository"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		userRepo := repository.NewUserRepository(db)
		userHandler := &handlers.UserHandler{Repo: userRepo}
		accountHandler := &handlers.AccountHandler{Repo: userRepo}
		r.Get("/users/{id}", userHandler.HandleGetUser)
		r.Post("/create-users/", userHandler.HandleCreateUser)
		r.Get("/accounts/{id}", accountHandler.HandleGetAccount)
		r.Get("/accounts/", accountHandler.HandleSomeDefaultHandler)
	})
	return r
}
