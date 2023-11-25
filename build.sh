#!/bin/bash

# Build Docker image
docker build -t victorbiga/trading-journal .

# Check if the build was successful
if [ $? -eq 0 ]; then
  # Run Docker container
  docker run -it --rm -p 8080:8080 victorbiga/trading-journal:latest
else
  echo "Error: Docker build failed."
fi
