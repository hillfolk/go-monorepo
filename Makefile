show:
	echo	"go-monorepo"

build:
	echo "Building..."
	go build -o bin/app-one ./apps/app-one/main.go
	go build -o bin/app-two ./apps/app-two/main.go

run-one:
	echo "Running... app-one"
	go run ./apps/app-one/main.go

run-two:
	echo "Running... app-two"
	go run ./apps/app-two/main.go


test:
	echo "Testing...";
	go test ./...;

coverage:
	echo "Coverage..."
	go test -coverprofile=coverage.out ./...;
	go tool cover -html=coverage.out;


all: build run-one run-two test coverage