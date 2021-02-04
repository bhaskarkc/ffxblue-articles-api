#!/usr/bin/env bash
#
# Docker container entrypoint script.

# set -e

APP_ENV=${APP_ENV:-local}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment..."

DSN="mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?multiStatements=true&charset=utf8&timeout=5s&autocommit=true"

echo "[`date`] Running DB migrations..."

migrate -database "${DSN}" -path ./migrations up

echo "[`date`] Starting server..."

ln -sf /dev/stdout /var/log/app/server.log

api-server >> /var/log/app/server.log 2>&1
