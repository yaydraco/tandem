# Contributing to Tandem

## Development Setup

### Requirements

- **Go >=1.24.2**: [Download](https://golang.org/dl/)
- **Docker**: [Install](https://docs.docker.com/get-docker/)

### Setup Instructions

1. Clone the repository:
   ```bash
   git clone https://github.com/yaydraco/tandem.git
   cd tandem
   ```

**For Nix/NixOS users:**
```bash
nix develop
```

**For everyone else:**

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   ```bash
   cp .example.env .env
   # Edit .env with your AI provider API keys
   ```

4. Build and run:
   ```bash
   go build -o tandem
   ./tandem
   ```