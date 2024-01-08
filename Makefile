postgres:
	docker run --name postgresLatest -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it postgresLatest createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresLatest dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/Kawaeugtkp/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock