# Run the CLI
run:
	go run main.go

# Build binary
build:
	go build -o devopscli main.go

# Install locally
install:
	go install

# Clean
clean:
	go clean -cache -modcache -testcache -fuzzcache
	rm -f devopscli

# Format the code
fmt:
	go fmt ./...

# Run with debug mode enabled
run-debug:
	go run main.go --debug=true

