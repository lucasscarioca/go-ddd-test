include .env
export

build:
	@go build -o bin/music-stash

run: build
	@./bin/music-stash

test:
	@go test -v ./...

create_migration:
	migrate create -ext sql -dir configs/migrations/ -seq $(name)

migration_up:
	migrate -path configs/migrations/ -database ${DB_URL} -verbose up

migration_down:
	migrate -path configs/migrations/ -database ${DB_URL} -verbose down

migration_fix:
	migrate -path configs/migrations/ -database ${DB_URL} force $(version)
