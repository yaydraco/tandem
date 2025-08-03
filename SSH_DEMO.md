# SSH Demo for Tandem

This directory contains a demonstration configuration for testing tandem's SSH functionality.

## Files:
- `demo-ssh-config.json` - Minimal configuration for SSH testing
- `RoE.md` - Sample Rules of Engagement file

## Usage:

Copy the demo config to the tandem config location:
```bash
mkdir -p .tandem
cp demo-ssh-config.json .tandem/swarm.json
```

Start the SSH server:
```bash
./tandem ssh --host localhost --port 2222
```

Connect from another terminal:
```bash
ssh localhost -p 2222
```

Note: This demo uses placeholder API keys and is meant for testing the SSH interface, not for actual penetration testing.