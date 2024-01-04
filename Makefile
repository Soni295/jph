dev:
	go run ./src/main.go

build:
	go build -o ./bin/main ./src/main.go

run:
	./bin/main

test:
	go test -v ./...