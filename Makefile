run:
	@go run main.go wire_gen.go

build:
	@go build -o chat main.go wire_gen.go

fix:
	@go vet .
	@gofmt -d -w .

