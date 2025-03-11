postgres: 
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:17.4-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root blog_post

dropdb:
	docker exec -it postgres17 dropdb blog_post

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/blog_post?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/blog_post?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc