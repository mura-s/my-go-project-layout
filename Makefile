.PHONY: build
build:
	CGO_ENABLED=0 go build -o ./bin/myapp ./cmd/myapp

.PHONY: test
test:
	go test ./...

.PHONY: coverage
coverage:
	go test -race -v -coverpkg=./... -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: deps
deps:
	go mod tidy
	go mod vendor

.PHONY: fmt
fmt:
	gofmt -s -w ./

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: clean
clean:
	rm ./bin/myapp
