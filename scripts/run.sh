#!/bin/bash

# Run a Go project by path (relative to root)
if [ -z "$1" ]; then
  echo "Usage: ./scripts/run.sh <project-path>"
  echo "Example: ./scripts/run.sh concepts/concurrency"
  exit 1
fi

cd "$1" || exit
go run .
