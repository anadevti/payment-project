package repository

import (
	"context"
	"database/sql"

	"github.com/anadevti/payment-project/internal/models"
	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user models.Users) error {
	user.ID = uuid.New()
	_, err := r.DB.ExecContext(ctx,
		"INSERT INTO users (id, name, email, cpf, phone, address, birth_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		user.ID, user.Name, user.Email, user.Cpf, user.Phone, user.Address, user.BirthDate, user.CreatedAt, user.UpdatedAt,
	)
	return err
}
