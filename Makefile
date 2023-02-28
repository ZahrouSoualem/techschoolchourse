postgres:
	docker run  --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.2

createdb:

	docker exec -it postgres12 createdb --username=root --owner=root simple_bank2

dropdb:

	docker exec -it postgres12 dropdb simple_bank2

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank2?sslmode=disable" -verbose up
	
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank2?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank2?sslmode=disable" -verbose up 1
	
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank2?sslmode=disable" -verbose down 1


test:
	go test -v -cover ./...

sqlc:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

server:
	go run main.go

mock:
	mockgen -package mockdb -destination  db/mock/store.go  tutorial.sqlc.dev/app/db/sqlc Store

.PHONY: postgres createdb dropdb sqlc migrateup migratedown migrateup1 migratedown1 test server mock