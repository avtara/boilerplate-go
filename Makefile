.PHONY: migrate-database
migrate-database:
	@migrate -path database/migrate -database "postgres://username:password@localhost:5432/boilerplate_go_db?sslmode=disable" up 2

