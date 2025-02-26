# 🌟 Stellar Service Manager (svcmgr)

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.20-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> Cosmic-grade CLI for managing server services with style 🚀

## ✨ Features

- **Unified Command Interface**
  - Add, list, login, remove, tunnel and backup services
  - Clean, intuitive CLI design with cosmic ASCII art
  - Modular command architecture using Cobra

- **Secure Operations** 
  - AES-256 encryption for sensitive data
  - Automatic key management in ~/.config/svcmgr/
  - OTP-based authentication support

- **Powerful Tunneling**
  - Easy SSH tunnel creation
  - Configurable local and remote ports
  - Connection tracking and monitoring

- **Enterprise Ready**
  - Structured JSON logging
  - Audit trail capabilities 
  - Cross-platform support

## 🚀 Quick Start

### Installation

```bash
# Clone repository
git clone https://github.com/moon9t/svcmgr.git

# Build binary
cd svcmgr && go build

# Run help
./svcmgr -h
```

## Basic Usage

```sh
# List services
svcmgr list

# Create tunnel
svcmgr tunnel myapp -l 8080 -r localhost -p 80

# Backup config
svcmgr backup

# Show version
svcmgr version
```

## 📖 Command Reference

### add
Add new cosmic service

```sh
svcmgr add
```

### list
Display services table

```sh
svcmgr list [--group-by type] [--show-hidden]
```

### login
Connect to service

```sh
svcmgr login [NAME]
```

### tunnel
Create secure connection

```sh
svcmgr tunnel [SERVICE] -l LOCAL_PORT -r REMOTE_HOST -p REMOTE_PORT
```

### backup
Store configuration securely

```sh
svcmgr backup
```

### version
Show build information

```sh
svcmgr version
```

## 🔧 Configuration

### Environment Variables
```sh
SVCMGR_DEBUG=true       # Enable debug logging
SVCMGR_CONFIG=path      # Config file location
```

### Directory Structure
```plaintext
~/.config/svcmgr/
  ├── vault.key         # Encryption key
  ├── config.yaml       # Main config
  └── services/         # Service definitions
```

## 🛠 Development

### Requirements

- Go 1.20+
- Git

### Build Process

1. **Get dependencies**
    ```sh
    go mod tidy
    ```

2. **Build binary**
    ```sh
    go build -o svcmgr
    ```

3. **Run tests**
    ```sh
    go test ./...
    ```

## Project Structure

```plaintext
.
├── cmd/              # Command line interface
├── internal/         # Internal packages
│   ├── config/       # Configuration management
│   ├── services/     # Service definitions
│   └── utils/        # Utility functions
└── main.go           # Entry point of the application
```

## 🤝 Contributing

Fork repository
Create feature branch
Commit changes
Push to branch
Open pull request

## 📜 License
MIT License - see LICENSE

Built with ❤️ by Moon9t | Where Code Meets Cosmos 🌠

