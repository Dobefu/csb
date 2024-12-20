#!/usr/bin/env bash

# Exit on error.
set -e

# Get the directory where the current script is located.
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)

# Navigate to the root of the project.
cd "$SCRIPT_DIR/.."

# Run the tests.
go test "./..." -cover -covermode=count
go tool cover -func coverage.out
