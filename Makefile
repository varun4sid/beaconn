.PHONY:	run build launch fmt vet test

run: fmt
	@go run ./cmd/beaconn/main.go

build: fmt vet test
	@go build -o ./build/beaconn ./cmd/beaconn/main.go
	@chmod +x ./build/beaconn

launch: build
	@./build/beaconn

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test ./...
