.PHONY: build
build:
	go build -o build/pr cmd/pr/main.go

test:
	go test -v ./...

run:
	./build/pr create --github-token XXX --branch test-branch --message "This is the commit message"

lint:
	golangci-lint run

generate:
	go generate ./...
