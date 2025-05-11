.PHONY: run build test clean

run:
	go run cmd/main.go

build:
	go build -o bin cmd/main.go

test:
	go test ./...

clean:
	rm -rf bin/