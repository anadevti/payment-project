package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
Recebe requisição HTTP
Valida entrada (body, params)
Chama serviço/repository
Retorna resposta JSON
Estuda: validação, status codes, error handling
*/

func GetAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("account id: %s", id)))
}

func SomeDefaultHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("id is required %s", id)))
		return
	}
}
