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

## Project Structure

This is a overview directory structure:

## **📌 Project Structure**

```
├── README.md
├── cmd/
│   ├── root.go
│   ├── root_test.go
│   └── version.go
├── config/
│   ├── config.go
│   ├── config_test.go
├── config.yaml
├── go.mod
├── go.sum
├── internal/
│   └── logger/
│       ├── logger.go
│       ├── logger_test.go
├── justfile
├── main.go
```

To bootstrap your directory with these:

```bash
mkdir -p {cmd,config,internal/logger}
touch cmd/{root,root_test,version}.go config/{config,config_test}.go internal/logger/{logger,logger_test}.go main.go justfile config.yaml
```

<details>
  <summary>The documentation</summary>

## **📖 Documentation**

### **2️⃣ Directory Breakdown**

#### **🗂 `cmd/` – CLI Commands**

This directory contains all **Cobra commands** for the CLI.

- **`root.go`**  
  - The **entry point** for the CLI.
  - Defines global flags and initializes Cobra commands.
  - Calls `Execute()` to start the CLI.
  
- **`root_test.go`**  
  - A **basic test** ensuring that the root command is set correctly.

- **`version.go`**  
  - Implements a `version` command.
  - Uses `viper.GetString("version")` to return the current CLI version.

#### **🗂 `config/` – Configuration Management**

Handles application **configuration settings** via **Viper**.

- **`config.go`**  
  - Loads configurations from `config.yaml`.
  - Provides defaults if no config file is found.
  - Enables `debug` mode if configured.

- **`config_test.go`**  
  - Tests `InitConfig()` to ensure proper config loading.

- **`config.yaml`** *(Optional)*  
  - Stores **default application settings**.
  - Example:
    ```yaml
    debug: false
    version: "1.0.0"
    ```

#### **🗂 `internal/logger/` – Logging Utilities**

A **helper package** for structured logging.

- **`logger.go`**  
  - Provides a `Log()` function that prints messages **only if debug mode is enabled**.

- **`logger_test.go`**  
  - Tests the logger to ensure **debug messages print correctly**.

#### **🗂 Root Files**

These are **core project files**.

- **`main.go`**  
  - The **entry point** of the application.
  - Calls `cmd.Execute()` to start the CLI.

- **`go.mod` / `go.sum`**  
  - **Dependency management** files for Go modules.

- **`justfile`**  
  - Defines **automation tasks** (if using `just`).
  - Example tasks:
    ```just
    build:
        go build -o starter main.go

    run:
        go run main.go
    ```

- **`README.md`**  
  - The **project documentation** file.

## **🎯 Usage**

### **🔹 Install Dependencies**

```sh
go mod tidy
```

### **🔹 Build the CLI**

```sh
go build -o starter main.go
```

### **🔹 Run the CLI**

```sh
./starter --help
```

### **🔹 Show Version**

```sh
./starter version
```

</details>

## Installing Dependencies

```bash
go mod init github.com/ruanbekker/go-cli-starter
go get github.com/spf13/cobra@latest
go get github.com/spf13/viper@latest
```

And installing from file:

```bash
go clean -cache -modcache -testcache -fuzzcache
go mod tidy
```

## Running the CLI

```bash
go run main.go --help          # See available commands
go run main.go version         # Display version
go run main.go version --debug # Display version with debug
```

Install as a binary:

```bash
go build -o starter main.go
```

Then run the example (it retrieves the version from config in `cmd/version.go`):

```bash
./starter version
# CLI Starter v1.0.0
```

## Extending to this example

This guide will help you extend the CLI by adding commands, handling configuration with Viper, and following best practices.

<details>
  <summary>Read more</summary>

## **Table of Contents**

1. [Adding a New Command](#adding-a-new-command)
2. [Using Configuration with Viper](#using-configuration-with-viper)
3. [Handling Logging & Debug Mode](#handling-logging--debug-mode)
4. [Best Practices for Extending the CLI](#best-practices-for-extending-the-cli)
5. [Automating Tasks with Just](#automating-tasks-with-just)
6. [Building & Distributing the CLI](#building--distributing-the-cli)

## **1. Adding a New Command**

To add a new command, follow these steps:

### **Step 1: Create a New Command File**

Navigate to the `cmd/` directory and create a new file for the command, e.g., `hello.go`:

```sh
touch cmd/hello.go
```

### **Step 2: Define the Command in `hello.go`**

```go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a greeting message",
	Long:  `The "hello" command prints a customizable greeting message.`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("debug") {
			fmt.Println("[DEBUG] Running hello command")
		}
		fmt.Println("Hello, Go CLI!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
```

### **Step 3: Run the New Command**

Rebuild and run the CLI:

```sh
go run main.go hello
```

**Output:**

```
Hello, Go CLI!
```

To run with debug mode:

```sh
go run main.go hello --debug
```

**Output:**

```
[DEBUG] Running hello command
Hello, Go CLI!
```

## **2. Using Configuration with Viper**

You can add new configuration options to `config/config.go`:

### **Example: Add a Configurable Greeting Message**

1. Update `config/config.go` to add a default message:

```go
viper.SetDefault("greeting", "Hello from the CLI!")
```

2. Modify `hello.go` to use this configuration:

```go
Run: func(cmd *cobra.Command, args []string) {
    if viper.GetBool("debug") {
        fmt.Println("[DEBUG] Running hello command")
    }
    fmt.Println(viper.GetString("greeting"))
},
```

3. Add a `config.yaml` file:

```yaml
greeting: "Hello, Gophers!"
```

4. Run:

```sh
go run main.go hello
```

**Output:**

```
Hello, Gophers!
```

## **3. Handling Logging & Debug Mode**

Debug mode can be enabled via:

```sh
go run main.go --debug
```

To add more structured logging, create a **logger utility**:

### **`internal/logger.go`**

```go
package internal

import (
	"log"
	"github.com/spf13/viper"
)

// Log prints messages only when debug mode is enabled
func Log(message string) {
	if viper.GetBool("debug") {
		log.Println("[DEBUG]", message)
	}
}
```

### **Use in Commands**

```go
import "github.com/yourusername/go-cli-starter/internal"

internal.Log("This is a debug message")
```

## **4. Best Practices for Extending the CLI**

- **Follow the Project Structure**:
  - Place new commands in `cmd/`
  - Store reusable utilities in `internal/`
  - Manage configuration in `config/`
  
- **Use Persistent Flags for Global Options**:

```go
rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose logging")
viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
```

- **Keep Commands Self-Contained**:
  - Each command should handle only one function.
  - Avoid putting too much logic in `cmd/root.go`.

## **5. Automating Tasks with Just**

The `justfile` simplifies CLI development tasks.

### **Add New Tasks**

Edit `justfile`:

```just
# Run with debugging enabled
run-debug:
    go run main.go --debug

# Test the new command
test-hello:
    go run main.go hello
```

Run:

```sh
just test-hello
```

## **6. Building & Distributing the CLI**

### **Build the Binary**

```sh
go build -o starter main.go
```

### **Install Locally**

```sh
go install
```

### **Distribute as a Binary**

Compile for multiple platforms:

```sh
GOOS=linux GOARCH=amd64 go build -o starter-linux main.go
GOOS=darwin GOARCH=arm64 go build -o starter-mac main.go
```

</details>
