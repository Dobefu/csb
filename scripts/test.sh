#!/usr/bin/env bash

# Exit on error.
set -e

# Get the directory where the current script is located.
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)

# Navigate to the root of the project.
cd "$SCRIPT_DIR/.."

# Run the tests.
go test "./..." -coverprofile="coverage.out" -covermode=count
go tool cover -func coverage.out
go tool cover -html=coverage.out
