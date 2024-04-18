include .env

build:
	go build -o ${BINARY} ./cmd/main.go

start:
	./${BINARY}

restart: build start