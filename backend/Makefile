.PHONY: help build run test clean lint docker-up docker-down migrate

help:
	@echo "SERGE Backend - Available commands:"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make lint         - Run go lint"
	@echo "  make vet          - Run go vet"
	@echo "  make fmt          - Format code"
	@echo "  make docker-up    - Start PostgreSQL container"
	@echo "  make docker-down  - Stop PostgreSQL container"
	@echo "  make migrate      - Run database migrations"

build:
	@echo "🔨 Building SERGE backend..."
	@go build -o bin/serge ./cmd

run: build
	@echo "🚀 Starting SERGE backend..."
	@./bin/serge

test:
	@echo "🧪 Running tests..."
	@go test -v ./...

clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf bin/
	@go clean

lint:
	@echo "🔍 Running golangci-lint..."
	@golangci-lint run ./...

vet:
	@echo "🔍 Running go vet..."
	@go vet ./...

fmt:
	@echo "📝 Formatting code..."
	@go fmt ./...

docker-up:
	@echo "🐘 Starting PostgreSQL container..."
	@docker-compose up -d
	@echo "✅ PostgreSQL is running on localhost:5432"

docker-down:
	@echo "🛑 Stopping PostgreSQL container..."
	@docker-compose down

migrate:
	@echo "🗄️  Running migrations..."
	@go run ./cmd migrate

dev: docker-up run
