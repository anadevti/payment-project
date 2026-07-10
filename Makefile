# Payment API - Makefile
# =====================

# Variáveis
APP_NAME=payment-api
MAIN=main.go
MIGRATIONS_DIR=db/migrations
DB_URL=postgres://payment_user:payment_password@localhost:5432/payment_db?sslmode=disable

# =====================
# Aplicação
# =====================

## Rodar a aplicação
run:
	go run $(MAIN)

## Build da aplicação
build:
	go build -o $(APP_NAME) $(MAIN)

## Limpar binário gerado
clean:
	rm -f $(APP_NAME)

## Instalar dependências
deps:
	go mod download
	go mod tidy

# =====================
# Docker
# =====================

## Subir PostgreSQL
docker-up:
	docker-compose up -d postgres

## Derrubar containers
docker-down:
	docker-compose down

## Ver logs do postgres
docker-logs:
	docker-compose logs -f postgres

## Resetar banco (apaga volume)
docker-reset:
	docker-compose down -v
	docker-compose up -d postgres

# =====================
# Migrations
# =====================

## Criar nova migration (uso: make migration NAME=create_users_table)
migration:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

## Rodar migrations (up)
migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

## Desfazer última migration (down)
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

## Ver versão atual da migration
migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

# =====================
# Testes
# =====================

## Rodar todos os testes
test:
	go test ./...

## Rodar testes com verbose
test-v:
	go test -v ./...

## Rodar testes com coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# =====================
# Utils
# =====================

## Ver todos os comandos disponíveis
help:
	@echo ""
	@echo "Payment API - Comandos disponíveis:"
	@echo "======================================"
	@grep -E '^##' Makefile | sed 's/## /  /'
	@echo ""

.PHONY: run build clean deps docker-up docker-down docker-logs docker-reset migration migrate-up migrate-down migrate-version test test-v test-coverage help