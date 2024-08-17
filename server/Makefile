postgresinit:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

# run the container and run the postgres sql command
postgres:
	docker exec -it postgres15 psql

# create datebase
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root go-chat

# drop the database without actually going into the container and doing it manually
dropdb:
	docker exec -it postgres15 dropdb go-chat

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/go-chat?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown

