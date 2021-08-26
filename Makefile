.PHONY: all clean test build fmt lint serve connect start list

ttt = tictactoe.bin

all: clean test build

clean:
	@echo "\nCleaning up..."
	rm -f ${ttt}
	@echo "Done."

test:
	@echo "\nTODO:..."
	@echo "\nTesting..."
	go test -v -coverpkg=./internal ./internal

build:
	@echo "\nTODO:..."
	@echo "\nBuilding..."
	@go version
	go build -o ${ttt}
	@echo "Done."

fmt format:
	go fmt ./...
	@echo "Done."

lint:
	golangci-lint run -v
	@echo "Done."

server serv serve:
	@echo "\nServing..."
	@go run ./cmd/server/main.go

client cli connect conn:
	@echo "\nConnecting..."
	@go run ./cmd/client/main.go

start: ${ttt}
	@echo "\nStarting..."
	./${ttt}

list:
	@grep '^[^#[:space:]].*:' Makefile

regen regenerate:
	@echo "\nRegenerating .proto..."
	./scripts/regen_proto.sh tictactoe.proto
	@echo "Done."