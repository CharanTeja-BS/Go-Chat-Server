# Go-Chat-Server

## Go Concurrent Port Scanner

A simple and fast TCP port scanner written in Go.  
Scans a specified range of ports on a target IP or domain concurrently and reports open ports.

---

## Features

- Scan TCP ports on any IP address or domain
- Specify port range with command line flags
- Uses goroutines for concurrent scanning (fast and efficient)
- Timeout handling for faster scanning
- Clean output showing only open ports

---

## Requirements

- Go 1.13 or higher

---

## Usage

```bash
go run main.go [--target=<IP or domain>] [--start=<start port>] [--end=<end port>]
