# devops-ai-cli

**devops-ai-cli** is a personal CLI tool written in **Go**, powered by **Cobra** and **Viper**. 

## About

It's a cli tool that includes some of the following:
- [Rendering Markdown](#-render-a-markdown-file)
- [Explain with OpenWebUI AI models from the terminal](#-explain-command)
- [Query with OpenWebUI - Continuing Conversations from the terminal](#-query-command-maintain-conversations)
- [Optimize Files: AI Recommendations](#-optimize-command)
- [Verify: Check if tools from config are installed](#-verify-installed-tools)

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

**📌 Option 2: Using Environment Variables**

```sh
export OPENWEB_API_KEY="your-secret-api-key"
export OPENWEB_API_HOST="http://localhost:3000"
```

#### **How It Works**

✅ The CLI sends a request to OpenWebUI with your query.  
✅ OpenWebUI processes the request and returns a response.  
✅ The response is rendered as **Markdown** using `glamour`.  

### **🤖 Query Command (Maintain Conversations)**

The `query` command allows you to **ask OpenWebUI AI questions** while maintaining **context** across multiple queries.

#### **📝 Basic Usage**

```sh
./devopscli query "What is Kubernetes?"
```

#### **Example Response**

```
  # Kubernetes Explained
  Kubernetes is an open-source container orchestration system...
🆔 **Conversation ID**: 1
```

_This assigns the question a **Conversation ID (`1`)**, allowing follow-ups._


#### **🔄 Continue a Conversation**

To ask **follow-up questions**, use the **`--cid`** flag with the **conversation ID**:

```sh
./devopscli query "How does Kubernetes handle scaling?" --cid "1"
```

#### **Example Response**

```
  # Kubernetes Scaling
  - Uses Horizontal Pod Autoscaler (HPA)
  - Uses Cluster Autoscaler
  - Supports Vertical Pod Autoscaling
```

_This maintains context and follows up on the previous question._

#### **📋 List Saved Conversations**

```sh
./devopscli query --list
```

#### **Example Output**

```
📝 **Previous Conversations:**
🆔 1: What is Kubernetes?
🆔 2: How does Kubernetes handle deployments?
```

_This lets you see which past questions you can continue._

#### **🗑️ Delete a Specific Conversation**

```sh
./devopscli query --delete 1
```

✅ Removes **conversation ID `1`** from storage.

#### **🚨 Clear All Conversations**

```sh
./devopscli query --clear
```

🗑️ **Deletes all stored conversations**.

#### **✨ Features**

✅ **Maintains conversation history**  
✅ **Allows follow-up questions (`--cid`)**  
✅ **Lists previous queries (`--list`)**  
✅ **Deletes single (`--delete`) or all (`--clear`) conversations**  
✅ **Uses OpenWebUI API for intelligent responses**  
✅ **Outputs beautifully formatted Markdown responses**  

### **🚀 Optimize Command**

The `optimize` command allows you to **send a code or configuration file** (e.g., **YAML, JSON, Python, Terraform, Shell scripts**) to **OpenWebUI AI**, which will analyze and provide **optimization suggestions in Markdown format**.

#### **🔹 Usage**

```sh
./devopscli optimize -f _extras/manifests/example-deployment.yaml 
```

**Example for a Terraform file:**
```sh
./devopscli optimize -f infra.tf
```

**Example for a Python script:**

```sh
./devopscli optimize -f script.py
```

### **⚙️ Configuration**

To use this command, configure OpenWebUI API details **via a config file or environment variables**.

#### **📌 Option 1: Using `config.yaml`**

```yaml
openwebui:
  host: "http://localhost:3000"
  api_key: "your-api-key-here"
  model: "gemma:2b"
```

#### **📌 Option 2: Using Environment Variables**

```sh
export OPENWEB_API_KEY="your-secret-api-key"
export OPENWEB_API_HOST="http://localhost:3000"
```

### **📝 Example Output**

If you run:

```sh
./devopscli optimize -f _extras/manifests/example-deployment.yaml
```

The response from AI will be returned in markdown format.

### **✨ Summary**

✅ **Supports multiple file types (YAML, JSON, .py, .tf, .sh, etc.)**  
✅ **Sends the file content to OpenWebUI for AI-based optimization**  
✅ **Receives Markdown suggestions and beautifully renders them in the terminal**  

### **🔍 Verify Installed Tools**

The `verify` command checks whether **required DevOps tools** are installed on your system. It reads the list of tools from **`config.yaml`** and reports their availability.

#### **📝 Usage**

```sh
./devopscli verify
```

#### **📌 Example Configuration**

Define required tools in **`config.yaml`**:

```yaml
tools:
  required:
    - kubectl
    - terraform
    - docker
    - helm
    - git
    - jq
    - curl
```

#### **📌 Example Output**

```sh
🔍 **Verifying Required Tools:**

✅ kubectl
❌ terraform (Not Installed)
✅ docker
✅ helm
✅ git
❌ jq (Not Installed)
✅ curl
```

#### **💡 Features**

✅ **Reads required tools from `config.yaml`**  
✅ **Checks if each tool is installed**  
✅ **Displays ✅ (installed) and ❌ (missing)**  
✅ **Fast and lightweight!**  


## 📚 Resources

- [Charmbracelet Glow](https://github.com/charmbracelet/glow)  
- [Charmbracelet Glamour](https://github.com/charmbracelet/glamour)  

