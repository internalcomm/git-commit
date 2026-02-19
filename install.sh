#!/usr/bin/env bash
set -e

APP=git-commit-linux
TMP_DIR="/tmp/$APP-install"
BASE_URL="https://github.com/internalcomm/git-commit/releases/latest/download"
ZIP_NAME="git-commit-dist.zip"

mkdir -p "$TMP_DIR"

echo "Downloading $APP..."
curl -fsSL "$BASE_URL/$ZIP_NAME" -o "$TMP_DIR/$ZIP_NAME"

echo "Extracting..."
unzip -o "$TMP_DIR/$ZIP_NAME" -d "$TMP_DIR"

chmod +x "$TMP_DIR/dist/$APP"
sudo mv "$TMP_DIR/dist/$APP" /usr/local/bin/git-commit

echo "✓ $APP installed successfully"
