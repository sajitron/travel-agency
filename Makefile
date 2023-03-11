DB_URL=postgresql://root:secret@localhost:5432/travel_agency?sslmode=disable

postgres:
	docker run --name postgres12 --network travel-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root travel_agency

dropdb:
	docker exec -it postgres12 dropdb travel_agency

startpg:
	docker start postgres12

stoppg:
	docker stop postgres12
	
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

dbdocs:
	dbdocs build doc/db.dbml

dbschema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/sajitron/travel-agency/db/sqlc Store

buildimage:
	docker build -t travel-agency:latest .

runcontainer:
	docker run --name travel-agency -p 2300:2300 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres12:5432/travel_agency?sslmode=disable" travel-agency:latest

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server startpg stoppg migrateup1 migratedown1 new_migration buildimage runcontainer dbschema mock