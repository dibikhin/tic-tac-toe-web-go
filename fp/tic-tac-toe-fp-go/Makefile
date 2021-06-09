.PHONY: all test lint run build start clean list

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

lint:
	golangci-lint run -v

run:
	@echo "\nRunning..."
	clear && go run main.go

start: ${ttt}
	@echo "\nStarting..."
	./${ttt}

list:
	@grep '^[^#[:space:]].*:' Makefile