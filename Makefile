postgres:
	docker run --name postgres_simplebank -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:15.6-alpine3.19

createdb:
	docker exec -it postgres_simplebank createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres_simplebank dropdb --username=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc