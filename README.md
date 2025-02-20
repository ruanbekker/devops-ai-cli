# devops-ai-cli

**devops-ai-cli** is a personal CLI tool written in **Go**, powered by **Cobra** and **Viper**. 

## About

It's a cli tool that includes some of the following:
- [Rendering Markdown](#-render-a-markdown-file)
- [Querying OpenWebUI AI models from the terminal](#-explain-command)

## 🚀 Installation

### **Prerequisites**

You need [Go](https://go.dev/dl/) installed. Check with:

```sh
go version
```

If Go is **not installed**, you can install version `1.24.0` on Linux

<details>
  <summary>Read more</summary>


```sh
curl -fsSL https://golang.org/dl/go1.24.0.linux-amd64.tar.gz | sudo tar -C /usr/local -xzf -
export PATH=$PATH:/usr/local/go/bin
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
echo "export GOPATH=\$HOME/go" >> ~/.bashrc
echo "export GOROOT=/usr/local/go" >> ~/.bashrc
```

Now, verify installation:

```sh
go version
```

</details>

## 🛠️ Setup & Usage

### **1️⃣ Install Dependencies**

```sh
go mod tidy
```

### **2️⃣ Build the CLI**

```sh
go build -o devopscli main.go
```

## ✨ Features

### **📜 Render a Markdown File**

The `render` command allows you to display Markdown files beautifully in the terminal.

#### **Usage**

```sh
./devopscli render -f example.md
```

#### **Example**

```sh
echo "# Hello, World!" > example.md
./devopscli render -f example.md
```

### **🤖 Explain Command**

The `explain` command allows you to ask **OpenWebUI AI** for explanations on various topics,  
and it returns a response **formatted in Markdown**.

#### **Usage**

```sh
./devopscli explain "what does the Kubernetes CrashLoopBackOff mean?"
```

#### **Configuration**

To use this command, configure OpenWebUI API details in one of two ways:

**📌 Option 1: Using `config.yaml`**

```yaml
openwebui:
  host: "http://localhost:3000"
  api_key: "your-api-key-here"
```

**📌 Option 2: Using an Environment Variable**

```sh
export OPENWEB_API_KEY="your-secret-api-key"
```

#### **How It Works**

✅ The CLI sends a request to OpenWebUI with your query.  
✅ OpenWebUI processes the request and returns a response.  
✅ The response is rendered as **Markdown** using `glamour`.  

## 📚 Resources

- [Charmbracelet Glow](https://github.com/charmbracelet/glow)  
- [Charmbracelet Glamour](https://github.com/charmbracelet/glamour)  

