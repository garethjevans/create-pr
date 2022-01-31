.PHONY: build
build:
	go build -o build/pr cmd/pr/main.go

test:
	go test ./...

run:
	./build/pr create --github-token XXX --branch branch-name --message "This is the commit message"

lint:
	golangci-lint run
