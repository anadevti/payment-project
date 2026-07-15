package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anadevti/payment-project/internal/models"
	"github.com/anadevti/payment-project/internal/repository"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func (h *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(422)
		w.Write([]byte("error: missing user id"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("user id: %s", id)))
}

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.Users

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Repo.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully!",
	})
}
