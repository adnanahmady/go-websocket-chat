run:
	@go run main.go wire_gen.go

build:
	@go build -o chat main.go wire_gen.go

fix:
	@go vet .
	@gofmt -d -w .

test:
	@go clean -testcache
	@go test -timeout 5s -failfast ./...
