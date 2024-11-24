#!/bin/bash
echo "Building Docker image..."
docker build -t transaction-summary .

echo "Running Docker container..."
docker run --rm transaction-summary