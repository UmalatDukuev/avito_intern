#!/bin/bash
set -e

echo "Waiting for database..."
sleep 5

echo "Running migrations..."
migrate -path /app/schema/migrations -database "$DATABASE_URL" up

echo "Starting app..."
exec ./main
