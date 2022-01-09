install-deps:
	@go mod download

swag:
	@swag init -g cmd/main.go

run: swag
	@docker-compose up --build

fmt:
	@go fmt ./...
