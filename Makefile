postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=vivek -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root banking_system

dropdb:
	docker exec -it postgres dropdb banking_system

migrateup:
	migrate -path db/migration -database "postgresql://root:vivek@localhost:5432/banking_system?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:vivek@localhost:5432/banking_system?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc