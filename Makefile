.PHONY:	run

run:
	@go run ./cmd/beaconn/main.go

build:
	@go build -o beaconn ./cmd/beaconn/main.go
	@chmod +x beaconn