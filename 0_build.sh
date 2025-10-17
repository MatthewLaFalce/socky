#!/bin/bash

set -e

APP_NAME="socky"
OUTPUT_DIR="./dist"

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[1;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

printf "${BLUE}🔧 Building %s for Linux (amd64)...${NC}\n" "$APP_NAME"

mkdir -p "$OUTPUT_DIR"

if GOOS=linux GOARCH=amd64 go build -o "${OUTPUT_DIR}/${APP_NAME}-linux" main.go; then
  printf "${GREEN}✅ Build complete: %s/${APP_NAME}-linux${NC}\n" "$OUTPUT_DIR"
else
  printf "${RED}❌ Build failed for Linux.${NC}\n"
  exit 1
fi

printf "${BLUE}🔧 Building %s for macOS (amd64)...${NC}\n" "$APP_NAME"

if GOOS=darwin GOARCH=amd64 go build -o "${OUTPUT_DIR}/${APP_NAME}-mac" main.go; then
  printf "${GREEN}✅ Build complete: %s/${APP_NAME}-mac${NC}\n" "$OUTPUT_DIR"
else
  printf "${RED}❌ Build failed for macOS.${NC}\n"
  exit 1
fi

