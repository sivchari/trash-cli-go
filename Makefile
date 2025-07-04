.PHONY: all build test lint clean install benchmark

BINARY_NAME=trash
GO=go

all: test build

build:
	$(GO) build -o $(BINARY_NAME) .

test:
	$(GO) test -v -race ./...

clean:
	rm -f $(BINARY_NAME)
	rm -rf benchmark/testdata
	rm -f benchmark/benchmark_results.txt

install: build
	cp $(BINARY_NAME) /usr/local/bin/

benchmark: build
	@echo "Creating test data..."
	@chmod +x benchmark/create_testdata.sh
	@./benchmark/create_testdata.sh
	@echo "Running benchmark..."
	@chmod +x benchmark/benchmark.sh
	@./benchmark/benchmark.sh

coverage:
	$(GO) test -v -race -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
