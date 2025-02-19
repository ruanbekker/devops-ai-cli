# go-cli-starter

## About

This is a **Go CLI Starter** project built with **Cobra** and **Viper**. It provides a modular, extensible structure for creating CLI tools.

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

Create a markdown file:

```bash
echo "# Hello World!" > example.md
```

Render markdown from a file:

```bash
./devopscli render -f example.md
```

## Resources

- [charmbracelet/glow](https://github.com/charmbracelet/glow)

