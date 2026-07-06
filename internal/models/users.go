package models

import (
	"github.com/google/uuid"
)

type Users struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Cpf     string    `json:"cpf"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
}
