#!/bin/bash

# Build the Docker image
docker build -t delivery-optimiser .

# Ensure output directory exists on the host
mkdir -p output

# Run the container with proper volume mappings
docker run --rm \
    -v $(pwd)/output:/app/output \
    -v $(pwd)/input.json:/app/input.json \
    -v $(pwd)/input_validation.json:/app/input_validation.json \
    delivery-optimiser

echo "Execution completed. Check 'output/' directory for results."