test:
	go test -v -cover ./...

coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out	

server:
	go run cmd/server/main.go

.PHONY: test coverage server