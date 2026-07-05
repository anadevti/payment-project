package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("erro ao abrir conexão:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}

	fmt.Println("✅ Conectado ao banco!")
	return db
}
