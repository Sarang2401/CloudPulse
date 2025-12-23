# CloudPulse

**Concurrent API Monitoring CLI built with Go**

CloudPulse is a lightweight DevOps-style command-line tool written in Go that concurrently checks the availability and latency of multiple APIs. It compiles into a single static binary, runs cross-platform, and demonstrates why Go is the language of choice for infrastructure and DevOps tooling.

---

## 🚀 Why This Project Exists
While learning Go, I wanted to understand why tools like Docker, Terraform, and Kubernetes are written in it. Instead of building a web app, I built a systems-focused CLI that mirrors real DevOps use cases:

* **Concurrent network calls**
* **Predictable resource usage**
* **Simple deployment model**

## ✨ Features
* **Concurrent health checks** for multiple APIs.
* **Measures HTTP status codes** and response latency.
* **Config-driven targets** via JSON.
* **Single static binary output** for easy distribution.
* **Cross-platform builds** (Linux / macOS / Windows).
* **No external dependencies.**

## 🧠 What This Demonstrates

| Concept | Implementation |
| :--- | :--- |
| **Concurrency** | Goroutines + channels |
| **Reliability** | Timeouts and explicit error handling |
| **Portability** | Static binary compilation |
| **DevOps Mindset** | CLI-first, config-driven design |

---

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

```text

--- 

⚙️ Configuration
config/targets.json

JSON

{
  "targets": [
    "https://api.github.com",
    "https://api.coindesk.com/v1/bpi/currentprice.json",
    "https://httpstat.us/200",
    "https://httpstat.us/503"
  ]
}
▶️ Running the Tool
Run locally
Bash

go run main.go
Sample Output
Plaintext

✅ https://api.github.com | Status: 200 | Latency: 140ms
❌ https://httpstat.us/503 | ERROR: 503 Service Unavailable
Note: All checks are executed concurrently, not sequentially.

📦 Building a Static Binary
Linux

Bash

GOOS=linux GOARCH=amd64 go build -o cloudpulse
macOS (Apple Silicon)

Bash

GOOS=darwin GOARCH=arm64 go build -o cloudpulse
Windows

Bash

GOOS=windows GOARCH=amd64 go build -o cloudpulse.exe
The resulting binary requires no Go runtime, has no external dependencies, and can be dropped directly onto a server or VM.

🔮 Future Enhancements
[ ] JSON output (--json) for automation.

[ ] Prometheus metrics export.

[ ] Retry and alerting logic.

[ ] Slack / webhook notifications.

[ ] Daemon mode for continuous monitoring.

🏁 Takeaway
This project helped me understand why Go is preferred for DevOps tools: predictable behavior, low overhead, easy concurrency, and simple deployment.