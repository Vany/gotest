VERSION := `git tag || git rev-parse --verify HEAD`
BIN := ${GOPATH}/bin


all: tidy main

main: main.go
	go build -o main .

test:
	go test .

tidy:
	go mod tidy

lint:
	${BIN}/golangci-lint run .

deps:
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
