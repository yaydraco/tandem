# Contributing to Tandem

Thank you for your interest in contributing to Tandem! This guide will help you get started with the development environment and contribute effectively to the project.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Development Environment Setup](#development-environment-setup)
- [Building the Project](#building-the-project)
- [Running the Project](#running-the-project)

## Overview

Tandem is a terminal-based application that provides a swarm of AI agents to assist in penetration testing engagements. The project is built with Go and uses a Terminal User Interface (TUI) powered by Bubble Tea. It integrates with multiple AI providers and maintains engagement context through Rules of Engagement (RoE) files.

### Key Technologies

- **Go 1.24.2+**: Primary programming language
- **Nix Flake**: Development environment management
- **Bubble Tea**: Terminal User Interface framework
- **SQLC**: Type-safe SQL code generation
- **SQLite**: Database for session and message storage
- **Vagrant**: Vulnerable lab environment setup
- **Multiple AI Provider SDKs**: OpenAI, Anthropic, Gemini, Groq, etc.

## Prerequisites

Before you begin, ensure you have the following installed:

### Required

1. **Nix Package Manager** (recommended approach)
   - Install Nix: https://nixos.org/download
   - Enable flakes: `echo "experimental-features = nix-command flakes" >> ~/.config/nix/nix.conf`

### Alternative (if not using Nix)

1. **Go 1.24.2+**
   - Download from: https://golang.org/dl/
   - Verify: `go version`

2. **SQLC** (for database code generation)
   - Install: `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

### Optional

1. **Vagrant** (for testing environments)
   - Download from: https://www.vagrantup.com/downloads
   - Used for setting up vulnerable lab environments (Metasploitable3)

2. **Git** (for version control)
   - Most systems have this pre-installed

## Development Environment Setup

### Using Nix (Recommended)

The project uses Nix flakes to provide a consistent, reproducible development environment across all platforms.

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yaydraco/tandem.git
   cd tandem
   ```

2. **Set up environment variables**:
   ```bash
   cp .example.env .env
   ```
   
   Edit `.env` and add your API keys for the AI providers you want to use:
   ```bash
   GEMINI_API_KEY=your_gemini_api_key_here
   OPENAI_API_KEY=your_openai_api_key_here
   GROQ_API_KEY=your_groq_api_key_here
   ANTHROPIC_API_KEY=your_anthropic_api_key_here
   # Add other providers as needed
   ```

3. **Enter the development shell**:
   ```bash
   nix develop
   ```

   This will automatically set up:
   - Go compiler and tools
   - Go language server (gopls)
   - Go development tools
   - SQLC for database code generation
   - Vagrant for lab environments
   - Starship prompt configuration

### Manual Setup (Alternative)

If you prefer not to use Nix:

1. **Clone and setup**:
   ```bash
   git clone https://github.com/yaydraco/tandem.git
   cd tandem
   cp .example.env .env
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Install development tools**:
   ```bash
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
   go install golang.org/x/tools/gopls@latest
   ```

## Building the Project

### Standard Build

```bash
# Build the main binary
go build -o tandem

# Build all packages (useful for checking compilation)
go build ./...
```

### Development Build

```bash
# Build with verbose output (useful for debugging)
go build -v ./...

# Build and run in one command
go run main.go
```

### Cross-platform Build

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build -o tandem-linux

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o tandem-macos

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o tandem.exe
```

## Running the Project

### Prerequisites for Running

1. **Create configuration directory**:
   ```bash
   mkdir -p .tandem
   ```

2. **Set up API keys** (if not done already):
   Ensure your `.env` file contains the necessary API keys for the AI providers you want to use.

3. **Create Rules of Engagement** (optional but recommended):
   ```bash
   echo "# Rules of Engagement

   ## Scope
   - Target: test.example.com
   - Authorization: Development testing only

   ## Constraints
   - No production systems
   - Local testing environment only" > .tandem/RoE.md
   ```

### Running the Application

```bash
# Run directly
./tandem

# Or run with Go
go run main.go

# Run with debug logging (if supported)
DEBUG=1 ./tandem
```

The application will start the Terminal User Interface (TUI) where you can interact with the AI agent swarm.