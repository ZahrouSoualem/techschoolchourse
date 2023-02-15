postgres:
	docker run  --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.2

createdb:

	docker exec -it postgres12 createdb --username=root --owner=root simple_bank2

dropdb:

	docker exec -it postgres12 dropdb simple_bank2

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank2?sslmode=disable" -verbose up
	
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank2?sslmode=disable" -verbose down



.PHONY: postgres createdb dropdb sqlc migrateup migrateup