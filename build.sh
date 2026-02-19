#!/usr/bin/env bash
set -e

echo "Building binaries..."

mkdir -p dist

GOOS=linux   GOARCH=amd64 go build -o dist/git-commit-linux
GOOS=darwin  GOARCH=amd64 go build -o dist/git-commit-macos
GOOS=windows GOARCH=amd64 go build -o dist/git-commit-windows.exe

echo "Build complete"
ls -lh dist
