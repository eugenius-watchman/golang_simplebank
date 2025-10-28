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

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
		proto/*.proto
		statik -src=./doc/swagger -dest=./doc 

# Evans gRPC client commands
evans:
	evans --host localhost --port 9090 --reflection

evans-proto:
	evans --host localhost --port 9090 --path proto proto/service_simple_bank.proto

# Test gRPC services
test-grpc: evans

# Create user via gRPC (example)
grpc-create-user:
	@echo 'Creating user via gRPC...'
	@evans --host localhost --port 9090 --reflection --call CreateUser

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 migrateup2 migratedown2 db_docs db_schema sqlc test server mock proto evans