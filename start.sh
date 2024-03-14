#!/bin/bash
set -e

echo "DB_SOURCE: $DB_SOURCE"  # Debug output

echo "Run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Start the app"
exec "$@"