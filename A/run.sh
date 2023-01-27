#!/bin/bash

# Run our services in a funny way

# Run docker-compose.yml files
# Run rabbitMQ
cd rabbitmq || exit
cp .env.example .env
docker compose up -d

# Run MySQL
cd ../mysql || exit
cp .env.example .env
docker compose up -d

# Run Gateway
cd ../gateway || exit
./main &

# Run Order
cd ../order || exit
./main &