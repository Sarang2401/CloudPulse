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

