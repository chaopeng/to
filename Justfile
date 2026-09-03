# List available recipes
default:
    @just --list

# Build binary
build:
    go build -o bin/to ./main.go

# Lint code using golangci-lint
lint:
    golangci-lint run ./...

# Run tests
test:
    go test -v ./...

# Format Go code using goimports
fmt:
    goimports -w -local "github.com/chaopeng/to" .

# Check formatting without modifying files
fmt-check:
    @test -z "$($(go env GOPATH)/bin/goimports -local github.com/chaopeng/to -l . 2>/dev/null || goimports -local github.com/chaopeng/to -l .)" || (echo "Unformatted Go files found:" && goimports -local github.com/chaopeng/to -l . && exit 1)

# Run go mod tidy
tidy:
    go mod tidy

# Update Go dependencies
update-deps:
    go get -u ./...
    @just tidy

# Clean build artifacts
clean:
    rm -rf bin
