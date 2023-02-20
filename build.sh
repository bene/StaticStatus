#!/bin/sh

# Download Tailwind CLI
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64

# Build CSS
./tailwindcss-linux-x64 -i ./src/main.css -o ./static/main.css

# Build status page
go run ./src/main.go
