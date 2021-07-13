.PHONY: all test lint serve build start clean list

ttt = tictactoe.bin

all: clean test build

clean:
	@echo "\nCleaning up..."
	rm -f ${ttt}

test:
	@echo "\nTesting..."
	go test -v -coverpkg=./internal ./internal

build:
	@echo "\nBuilding..."
	@go version
	go build -o ${ttt}

fmt:
	go fmt ./...

lint:
	golangci-lint run -v

serve:
	@echo "\nRunning..."
	go run ./cmd/server/main.go

connect:
	@echo "\nRunning..."
	go run ./cmd/client/main.go

start: ${ttt}
	@echo "\nStarting..."
	./${ttt}

list:
	@grep '^[^#[:space:]].*:' Makefile