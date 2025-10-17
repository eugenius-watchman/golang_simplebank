postgres:
	docker run --name postgres13 --network bank-net -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres13 dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up 1

migrateup2:
	migrate -path db/migration -database "postgresql://root:JukrzNXlvDmOn9EsrZ4X@localhost:5432/simple_bank?sslmode=disable" -verbose up 2

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down 1

migratedown2:
	migrate -path db/migration -database "postgresql://root:JukrzNXlvDmOn9EsrZ4X@localhost:5432/simple_bank?sslmode=disable" -verbose down 2

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	 dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/eugenius-watchman/golang_simplebank/db/sqlc Store
	
.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 migrateup2 migratedown2 db_docs db_schema sqlc test server mock