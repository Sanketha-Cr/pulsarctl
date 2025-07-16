#!/bin/bash
# Build script for s390x architecture support

echo "Building pulsarctl for s390x architecture..."

# Set environment for cross-compilation
export GOOS=linux
export GOARCH=s390x

# Build the main binary
go build -o bin/pulsarctl-s390x ./

# Build specific packages that had issues
echo "Building topic package..."
go build ./pkg/ctl/topic/...

echo "Build completed for s390x architecture"