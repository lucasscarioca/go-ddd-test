build:
	@go build -o bin/music-stash

run: build
	@./bin/music-stash

test:
	@go test -v ./...
