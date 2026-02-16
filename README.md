# 🚀 Go Concurrent Site Monitor

A high-performance, concurrent health-check tool written in **Go (Golang)**. 
Designed to demonstrate **Goroutines**, **Channels**, and **Cloud-Native packaging**.

## 🛠 Tech Stack
* **Language:** Go (Golang)
* **Concurrency:** Goroutines & Channels
* **Deployment:** Docker / Containerized

## ⚡ How It Works
Unlike traditional sequential scripts that check websites one by one (blocking), this tool:
1. **Spawns a Goroutine** for every target URL instantly.
2. **Checks all sites concurrently**, drastically reducing total execution time.
3. **Aggregates results** via a thread-safe Channel.
4. **Handles Timeouts & Errors** gracefully (e.g., detecting down sites immediately).

## 📦 How to Run

### Option 1: Local Execution (Go Installed)
Ensure you have Go installed, then run:
```bash
go mod init go-site-monitor-system
go run main.go
```
1. go mod init go-site-monitor-system
The Meaning: We are creating the project's ID Card.

What it does: It tells Go, "This folder is now a legitimate software project named go-site-monitor-system."

The Result: It creates a file named go.mod. Without this, Go treats our files as just loose text scripts.

2. go run main.go
The Meaning: We are doing a Test Drive.

What it does: Go takes our code, instantly builds it into a working program in a hidden temporary folder, runs it once so we can see the output, and then cleans up.

Why use it: It is the fastest way to check if our code works before we build the final permanent version.

---

### **Option 2: The Dockerfile**

Now, let's make sure the `Dockerfile` is perfect. This file tells Docker how to package **our** app.

Paste this into **our** `Dockerfile` (make sure the filename is exactly `Dockerfile` with no extension):

```dockerfile
# 1. Start from a small, secure base image (Alpine Linux)
# We use the official Go image to compile our code
FROM golang:alpine

# 2. Set the working directory inside the container
WORKDIR /app

# 3. Copy our source code into the container
COPY . .

# 4. Build the application
# This creates a binary executable named "monitor"
RUN go build -o monitor main.go

# 5. Command to run the executable when the container starts
CMD ["./monitor"]
```

### **Finally Execusion**
# 1. Build the Docker Image
# This creates a lightweight Alpine Linux container with the compiled binary
docker build -t site-monitor .

# 2. Run the Container
docker run site-monitor

---

## 👤 Author
**Md Ahsanul Bari (Romi)**
