# devops-ai-cli

## About

This is a personal project to have a cli terminal tool written in Go with Viper and Cobra.

## Pre-requirements

You will need [Go](https://go.dev/dl/) installed, you can test using:

```bash
go version
```

<details>
  <summary>If that fails, read more</summary>

If you see: `go: command not found`, you will need to install Go. For version `1.24.0` on Linux, you can run:

```bash
curl -fsSL https://golang.org/dl/go1.24.0.linux-amd64.tar.gz | sudo tar -C /usr/local -xzf -
export PATH=$PATH:/usr/local/go/bin
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
echo "export GOPATH=\$HOME/go" >> ~/.bashrc
echo "export GOROOT=/usr/local/go" >> ~/.bashrc
```

Now you should get a version back when running `go version` and then to test:

```bash
echo 'package main; import "fmt"; func main() { fmt.Println("Hello, Go!") }' > test.go
go run test.go
rm -f test.go
```

</details>

## Steps

<details>
  <summary>View the steps</summary>

Download the `glow` library:

```bash
go get github.com/charmbracelet/glow
go get github.com/charmbracelet/glamour
```

</details>

## Features

### Render a Markdown File

Install dependencies:

```bash
go mod tidy
```

Build the cli:

```bash
go build -o devopscli main.go
```

Create a markdown file:

```bash
echo "# Hello World!" > example.md
```

Render markdown from a file:

```bash
./devopscli render -f example.md
```

### Explain command

The `explain` command allows you to ask OpenWebUI for explanations on various topics, and it will return a response formatted in Markdown.

#### Usage

```bash
./devopscli explain "what does the Kubernetes CrashLoopBackOff mean?"
```

#### Configuration

To use this command, you need to configure OpenWebUI API details. You can do this in one of two ways:

Option 1: Using config.yaml:

```yaml
openwebui:
  host: "http://localhost:3000"
  api_key: "your-api-key-here"
```

Option 2: Using an Environment Variable:

```bash
export OPENWEB_API_KEY="your-secret-api-key"
```

#### How it works

- The CLI sends a request to OpenWebUI with your query.
- OpenWebUI processes the request and returns a response.
- The response is rendered as Markdown using `glamour`.

## Resources

- [charmbracelet/glow](https://github.com/charmbracelet/glow)

