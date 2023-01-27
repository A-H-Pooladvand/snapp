#!/bin/bash

# Building go services
# Check if Go is installed
if ! [ -x "$(command -v go)" ]; then
  echo 'Go is not installed. Installing...'
  sudo apt-get update
  sudo apt-get install -y golang-go
else
  cd gateway
  # Check if .env file exists
  if [ ! -f ".env" ]; then
    # If not, copy it from .env.example
    cp .env.example .env
  fi
  go build -o main .
  chmod +x main

  # SHOULD BE INSIDE A FOR LOOP BUT WHATEVER...
    cd ../order
    # Check if .env file exists
    if [ ! -f ".env" ]; then
      # If not, copy it from .env.example
      cp .env.example .env
    fi
    go build -o main .
    chmod +x main
fi

