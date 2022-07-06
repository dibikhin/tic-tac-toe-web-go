.PHONY: fmt lint test serve connect regen list

fmt format:
	goimports -l -w .
	@echo "Done."

lint:
	golangci-lint run -v
	@echo "Done."

test:
	@echo "\nTesting client..."
	go test -v -coverpkg=./pkg/client ./pkg/client
	@echo "\nTesting server..."
	go test -v -coverpkg=./pkg/server ./pkg/server

server serv serve:
	@echo "\nServer:"
	@go run ./cmd/server/main.go

client cli connect conn:
	@echo "\nClient:"
	@go run ./cmd/client/main.go

regen regenerate:
	@echo "\nRegenerating .proto..."
	./scripts/regen_proto.sh tictactoe.proto
	@echo "Done."

list:
	@grep '^[^#[:space:]].*:' Makefile