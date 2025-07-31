# Contributing to Tandem

Thank you for your interest in contributing to Tandem! This guide will help you get started with the development environment and contribute effectively to the project.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Development Environment Setup](#development-environment-setup)
- [Building the Project](#building-the-project)
- [Running the Project](#running-the-project)
- [Testing](#testing)
- [Code Generation](#code-generation)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Contributing Workflow](#contributing-workflow)
- [Code Style and Standards](#code-style-and-standards)
- [Troubleshooting](#troubleshooting)

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

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Run tests for a specific package
go test ./internal/config
```

### Test Structure

The project currently has minimal test coverage. Tests are located alongside the source code in `*_test.go` files. The main test file is:

- `internal/config/config_test.go`: Tests for configuration and agent prompt generation

### Writing Tests

When adding new functionality, please include tests:

1. **Unit tests**: Test individual functions and methods
2. **Integration tests**: Test interactions between components
3. **Table-driven tests**: Use Go's table-driven test pattern for multiple test cases

Example test structure:
```go
func TestYourFunction(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        expected string
    }{
        {"case1", "input1", "expected1"},
        {"case2", "input2", "expected2"},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := YourFunction(tc.input)
            if result != tc.expected {
                t.Errorf("Expected %v, got %v", tc.expected, result)
            }
        })
    }
}
```

## Code Generation

The project uses SQLC for type-safe database code generation.

### SQLC Setup

SQLC configuration is defined in `sqlc.json`:
- **Schema location**: `internal/db/migration`
- **Query location**: `internal/db/sql`
- **Generated code**: `internal/db/`

### Regenerating Database Code

After modifying SQL queries or schema:

```bash
# Generate Go code from SQL
sqlc generate

# Verify generated code compiles
go build ./internal/db
```

### Database Migrations

Migration files are located in `internal/db/migrations/`. Follow the naming convention:
```
YYYYMMDDHHMMSS_description.sql
```

### SQL Queries

Write your SQL queries in `internal/db/sql/` and SQLC will generate type-safe Go code.

## Project Structure

```
tandem/
├── .github/                 # GitHub-specific files
├── .tandem/                 # Configuration directory
├── internal/                # Internal Go packages
│   ├── agent/              # AI agent management
│   ├── app/                # Application core
│   ├── cmd/                # CLI commands (Cobra)
│   ├── config/             # Configuration management
│   ├── db/                 # Database layer (SQLC generated)
│   ├── diff/               # Diff utilities
│   ├── format/             # Text formatting
│   ├── logging/            # Logging utilities
│   ├── message/            # Message handling
│   ├── models/             # Data models
│   ├── provider/           # AI provider integrations
│   ├── pubsub/             # Pub/Sub messaging
│   ├── session/            # Session management
│   ├── tools/              # Tool integrations
│   ├── tui/                # Terminal User Interface
│   ├── utils/              # General utilities
│   └── version/            # Version information
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── flake.nix               # Nix development environment
├── sqlc.json               # SQLC configuration
├── Vagrantfile             # Vagrant lab environment
├── .example.env            # Environment variables template
└── README.md               # Project documentation
```

### Key Directories

- **`internal/`**: Contains all internal Go packages following Go best practices
- **`internal/tui/`**: Terminal User Interface components using Bubble Tea
- **`internal/db/`**: Database layer with SQLC-generated code
- **`internal/agent/`**: AI agent management and coordination
- **`internal/provider/`**: Integrations with different AI providers

## Configuration

### Environment Variables

Copy `.example.env` to `.env` and configure:

```bash
# AI Provider API Keys
GEMINI_API_KEY=your_key_here
OPENAI_API_KEY=your_key_here
GROQ_API_KEY=your_key_here
OPENROUTER_API_KEY=your_key_here
VERTEX_API_KEY=your_key_here
XAI_API_KEY=your_key_here
ANTHROPIC_API_KEY=your_key_here
COPILOT_API_KEY=your_key_here
```

### Agent Configuration

Agents are configured in `.tandem/swarm.json`. This file defines:
- Agent roles and capabilities
- AI models to use for each agent
- Agent-specific instructions and goals
- Tool permissions

### Rules of Engagement

Create `.tandem/RoE.md` in your working directory to provide context for penetration testing engagements.

## Contributing Workflow

### 1. Setting Up for Contribution

1. Fork the repository on GitHub
2. Clone your fork locally
3. Set up the development environment (see above)
4. Create a new branch for your feature/fix

```bash
git checkout -b feature/your-feature-name
```

### 2. Making Changes

1. Make your changes following the code style guidelines
2. Add tests for new functionality
3. Ensure all tests pass
4. Update documentation if necessary

### 3. Testing Your Changes

```bash
# Run tests
go test ./...

# Build to ensure compilation
go build ./...

# Test the application manually
go run main.go
```

### 4. Submitting Changes

1. Commit your changes with clear, descriptive messages
2. Push to your fork
3. Create a Pull Request on GitHub
4. Ensure CI passes and address any feedback

```bash
git add .
git commit -m "feat: add new feature description"
git push origin feature/your-feature-name
```

### Commit Message Format

Use conventional commit format:
- `feat:` new features
- `fix:` bug fixes
- `docs:` documentation changes
- `style:` formatting changes
- `refactor:` code refactoring
- `test:` adding/updating tests
- `chore:` maintenance tasks

## Code Style and Standards

### Go Style Guidelines

1. **Follow Go conventions**:
   - Use `gofmt` for formatting
   - Follow Go naming conventions
   - Use Go modules for dependencies

2. **Code organization**:
   - Keep packages focused and cohesive
   - Use interfaces for abstraction
   - Handle errors explicitly

3. **Documentation**:
   - Add comments for exported functions
   - Use Go doc conventions
   - Keep comments concise but clear

### Project-Specific Guidelines

1. **Use `any` instead of `interface{}`** (per project preference)
2. **Stick to existing dependencies** - avoid adding new dependencies without discussion
3. **When implementing interfaces**, provide method signatures with empty bodies initially

Example interface implementation:
```go
// Generate a stub implementing BaseTool interface for DockerCli
func (d *DockerCli) Info() tools.ToolInfo {}
```

### Linting and Formatting

```bash
# Format code
go fmt ./...

# Run go vet
go vet ./...

# If you have golangci-lint installed
golangci-lint run
```

## Troubleshooting

### Common Issues

1. **Build failures due to missing dependencies**:
   ```bash
   go mod tidy
   go mod download
   ```

2. **SQLC generation issues**:
   ```bash
   # Ensure SQLC is installed
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
   
   # Regenerate
   sqlc generate
   ```

3. **Environment setup issues with Nix**:
   ```bash
   # Clean and rebuild
   nix flake update
   nix develop --refresh
   ```

4. **API key issues**:
   - Ensure `.env` file exists and contains valid API keys
   - Check that environment variables are properly loaded

### Getting Help

- **Issues**: Create an issue on GitHub for bugs or feature requests
- **Discussions**: Use GitHub Discussions for questions and general discussion
- **Documentation**: Check the README.md and project wiki

### Development Tips

1. **Use the Nix development shell** for consistency
2. **Test with different AI providers** if you have access
3. **Run tests frequently** during development
4. **Keep commits small and focused**
5. **Update documentation** alongside code changes

## Lab Environment (Optional)

For testing penetration testing scenarios, the project includes a Vagrantfile for setting up vulnerable lab environments:

```bash
# Start Metasploitable3 Ubuntu
vagrant up ub1404

# Start Metasploitable3 Windows
vagrant up win2k8

# SSH into Ubuntu environment
vagrant ssh ub1404
```

This provides realistic targets for testing the penetration testing agents.

---

Thank you for contributing to Tandem! Your contributions help make this tool better for the penetration testing community.