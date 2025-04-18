.PHONY: build run test clean lint

build:
	go build -o gator main.go

run:
	go run main.go

test:
	go test -v ./tests

lint:
	golangci-lint run ./...

clean:
	rm -f gator