init: 
	@go run github.com/99designs/gqlgen init

gen:
	@go run github.com/99designs/gqlgen generate

start:
	@go run ./cmd/server/server.go