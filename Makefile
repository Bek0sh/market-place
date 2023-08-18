.PHONY: createMigration migrateUp createDB rundb migrationUp migrationDown

createMigration:
	migrate create -ext sql -dir pkg/db/migrations -seq init_schema

migrationUp:
	migrate -path pkg/db/migrations -database "postgres://postgres:1234@localhost:5432/market?sslmode=disable" -verbose up

migrationDown:
	migrate -path pkg/db/migrations -database "postgres://postgres:1234@localhost:5432/market?sslmode=disable" -verbose down

createDB:
	docker exec -it postgres15 createdb --username=postgres market

rundb:
	docker start postgres15