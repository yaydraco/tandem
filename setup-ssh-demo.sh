#!/bin/bash

# SSH Demo Test Script for Tandem
# This script sets up and tests the SSH functionality

set -e

echo "🚀 Setting up tandem SSH demo..."

# Create demo directory structure
mkdir -p .tandem

# Copy demo configuration
echo "📋 Setting up demo configuration..."
cp demo-ssh-config.json .tandem/swarm.json
cp demo-RoE.md .tandem/RoE.md

echo "✅ Demo configuration ready!"
echo ""
echo "📚 To test the SSH functionality:"
echo ""
echo "1. Start the SSH server:"
echo "   ./tandem ssh --host localhost --port 2222"
echo ""
echo "2. In another terminal, connect via SSH:"
echo "   ssh localhost -p 2222"
echo ""
echo "3. You should see the tandem TUI interface over SSH!"
echo ""
echo "🔧 SSH server configuration:"
echo "   - Host: localhost"
echo "   - Port: 2222" 
echo "   - Host key will be auto-generated at: .ssh/tandem_host_key"
echo ""
echo "💡 Note: This is a demo configuration with placeholder API keys."
echo "   The interface will work but agents won't make real API calls."
echo ""
echo "🎯 The SSH demo shows tandem's capability to serve its TUI"
echo "   over SSH, enabling remote multi-user access to penetration"
echo "   testing capabilities."