#!/usr/bin/env bash
#
# Docker container entrypoint script.

# set -e

# Logging
# i. Write stdout to entry.log
# ii.System's logger reads entry.log and writes log into system log tags
# iii. errors are redirected to stdout
exec > >(tee -a /var/log/app/entry.log|logger -t server -s 2>/dev/console) 2>&1

APP_ENV=${APP_ENV:-local}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment..."

DSN="mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?multiStatements=true&charset=utf8&timeout=5s&autocommit=true"

echo "[`date`] Running DB migrations..."

migrate -database "${DSN}" -path ./migrations up

echo "[`date`] Starting server..."

api-server >> /var/log/app/server.log 2>&1
