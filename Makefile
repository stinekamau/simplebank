postgresbank:
	docker run --name postgresbank -e  POSTGRESS_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:15-alpine

createdb: 
	docker exec -it postgresbank createdb -U root -O root bank 

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down

dropdb: 
	docker exec -it postgresbank dropdb  bank 



sqlc: 
	sqlc generate

test: 
	go test -v -cover ./...

server: 
	go run main.go 

.PHONY:postgresbank createdb  dropdb  migrateup migratedown sqlc test server


