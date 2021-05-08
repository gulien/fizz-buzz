.PHONY: fmt lint tests todos godoc run build

# Formats the code, "optimizes" the modules' dependencies.
fmt:
	go fmt ./...
	go mod tidy

# Runs linters.
lint:
	golangci-lint run

# Runs tests.
tests:
	go test -race -covermode=atomic -coverprofile=coverage.txt ./... &&\
	go tool cover -html=coverage.txt -o coverage.html

# Shows TODOs.
todos:
	golangci-lint run \
	--no-config \
	--disable-all \
	--enable godox

# Runs a local webserver for godoc.
godoc:
	$(info http://localhost:6060/pkg/github.com/gulien/fizz-buzz)
	godoc -http=:6060

# Runs the application.
PORT=80
TIMEOUT=30

run:
	go run cmd/fizzbuzz/main.go --port=$(PORT) --timeout=$(TIMEOUT)


# Builds the application.
VERSION=snapshot

build:
	go build -ldflags "-X main.version=$(VERSION)" -o=fizzbuzz cmd/fizzbuzz/main.go