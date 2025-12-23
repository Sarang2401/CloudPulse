# CloudPulse ⚡

**Concurrent API Monitoring CLI built with Go**

## 🚀 Why This Project Exists

CloudPulse is a lightweight, production-style command-line tool written in Go that concurrently checks the availability and latency of multiple APIs.

While learning Go, I wanted to understand *why tools like Docker, Terraform, and Kubernetes are written in Go*.  
Instead of building a web application, I built a **systems-focused CLI** that mirrors real DevOps use cases:

- Concurrent network calls  
- Predictable performance and resource usage  
- Simple, portable deployment via a single binary  

The goal is to demonstrate **engineering fundamentals**, not UI complexity.

## ✨ Features

- ⚡ Concurrent health checks for multiple APIs  
- ⏱ Measures HTTP status codes and latency  
- 📄 Config-driven targets using JSON  
- 📦 Single static binary (no runtime dependencies)  
- 🖥 Cross-platform builds (Linux / macOS / Windows)  
- 🧼 No external dependencies (standard library only)

## 🧠 What This Demonstrates

| Concept | Implementation |
|------|---------------|
| Concurrency | Goroutines + channels |
| Reliability | Timeouts and explicit error handling |
| Portability | Static binary compilation |
| DevOps Mindset | CLI-first, config-driven design |
| Observability Basics | Latency measurement & status reporting |

## 📁 Project Structure

```text
cloudpulse/
├── main.go                # CLI entry point
├── checker/
│   ├── check.go           # API check logic
│   └── result.go          # Result data structure
├── config/
│   └── targets.json       # API endpoints to monitor
├── go.mod
└── README.md


---

```md
## ⚙️ Configuration

Define the APIs you want to monitor in `config/targets.json`:

{
  "targets": [
    "https://api.github.com",
    "https://api.coindesk.com/v1/bpi/currentprice.json",
    "https://httpstat.us/200",
    "https://httpstat.us/503"
  ]
}

## ▶️ Running the Tool

Run locally:

go run main.go

Sample Output:

✅ https://api.github.com | Status: 200 | Latency: 140ms
❌ https://httpstat.us/503 | ERROR: 503 Service Unavailable

Note: All API checks are executed concurrently, not sequentially.

## 📦 Building a Static Binary

### Linux

GOOS=linux GOARCH=amd64 go build -o cloudpulse

### macOS (Apple Silicon)

GOOS=darwin GOARCH=arm64 go build -o cloudpulse

### Windows

GOOS=windows GOARCH=amd64 go build -o cloudpulse.exe

The resulting binary:
- Requires no Go runtime
- Has no external dependencies
- Can be deployed directly to servers or VMs

## 🔮 Future Enhancements

- [ ] JSON output (`--json`) for automation pipelines  
- [ ] Prometheus metrics export  
- [ ] Retry logic with backoff  
- [ ] Slack / webhook notifications  
- [ ] Daemon mode for continuous monitoring

## 🏁 Takeaway

CloudPulse demonstrates why Go is preferred for DevOps and infrastructure tooling:
predictable behavior, low overhead, native concurrency, and simple deployment.

This project reflects a **production-first engineering mindset** focused on correctness, portability, and clarity.
