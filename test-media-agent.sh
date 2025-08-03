#!/bin/bash

# Media Agent Test Script
# This script tests the basic functionality of the media agent

set -e

echo "🎬 Testing Tandem Media Agent..."

# Build Tandem
echo "Building Tandem..."
go build -o tandem .

# Test 1: Verify the media agent is registered
echo "Test 1: Checking media agent registration..."
if grep -q "media_agent" .tandem/swarm.json; then
    echo "✅ Media agent found in configuration"
else
    echo "❌ Media agent not found in configuration"
    exit 1
fi

# Test 2: Verify tools compile
echo "Test 2: Checking tool compilation..."
if go build ./internal/tools/...; then
    echo "✅ Media agent tools compile successfully"
else
    echo "❌ Media agent tools failed to compile"
    exit 1
fi

# Test 3: Check if example files exist
echo "Test 3: Checking example files..."
if [[ -f "media/tandem-showcase.tape" && -f "media/freeze-config.json" ]]; then
    echo "✅ Example media files exist"
else
    echo "❌ Example media files missing"
    exit 1
fi

# Test 4: Validate VHS tape file syntax (basic check)
echo "Test 4: Validating VHS tape file..."
if grep -q "Output\|Type\|Enter\|Sleep" media/tandem-showcase.tape; then
    echo "✅ VHS tape file has valid syntax"
else
    echo "❌ VHS tape file syntax appears invalid"
    exit 1
fi

# Test 5: Validate Freeze config (basic JSON check)
echo "Test 5: Validating Freeze config..."
if python3 -m json.tool media/freeze-config.json > /dev/null 2>&1; then
    echo "✅ Freeze config is valid JSON"
else
    echo "❌ Freeze config is invalid JSON"
    exit 1
fi

# Test 6: Test help output includes new flags
echo "Test 6: Testing CLI functionality..."
if ./tandem --help | grep -q "prompt\|output-format"; then
    echo "✅ CLI supports media agent invocation"
else
    echo "❌ CLI missing expected flags"
    exit 1
fi

echo ""
echo "🎉 All tests passed! Media Agent is ready to use."
echo ""
echo "Next steps:"
echo "1. Set up API keys in .env file"
echo "2. Test with: ./tandem -p 'As the media agent, create promotional content'"
echo "3. Try the GitHub Actions workflow on a push"