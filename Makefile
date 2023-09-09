build: 
	@go build -o bin/realTimeDb cmd/main.go

run: build
	@./bin/realTimeDb

test:
	@go test -v ./...