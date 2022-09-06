include app.env
export
run:
	@go run cmd/main.go
migrate_create:
	migrate create -ext sql -dir ./internal/storage/persistant/migration -seq gym
sqlc:
	sqlc generate