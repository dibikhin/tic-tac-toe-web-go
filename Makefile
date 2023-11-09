.PHONY: fmt lint test serve connect regen list

fmt format:
	goimports -l -w .
	@echo "Done."

lint:
	golangci-lint run -v
	@echo "Done."

test:
	@echo "\nTesting client..."
<<<<<<< Updated upstream
	go test -v -coverpkg=./pkg/gameclient ./pkg/gameclient
	@echo "\nTesting server..."
	go test -v -coverpkg=./pkg/gameserver ./pkg/gameserver
=======
	go test -v -coverpkg=./client ./client
	@echo "\nTesting server..."
	go test -v -coverpkg=./server ./server

ci:
	make fmt
	make lint
	make test
>>>>>>> Stashed changes

server serv serve:
	@echo "\nServer:"
	@go run ./cmd/server

client cli connect conn:
	@echo "\nClient:"
	@go run ./cmd/client

regen regenerate:
	@echo "\nRegenerating .proto..."
	./scripts/regen_proto.sh tictactoe.proto
	@echo "Done."

list:
	@grep '^[^#[:space:]].*:' Makefile