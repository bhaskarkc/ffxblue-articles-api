# Dockerfile for the ffxblue-articles-api
# Multistage Dockerfile
#
# Author: Bhaskar K <xlinkerz@gmail.com>

FROM golang:alpine AS build
RUN apk update && apk add curl git bash make ca-certificates && rm -rf /var/cache/apk/*

# golang-migrate for db migrations
# https://github.com/golang-migrate/migrate/releases
ARG MIGRATE_VERSION=4.14.1
ADD https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz /tmp
RUN tar -xzf /tmp/migrate.linux-amd64.tar.gz -C /usr/local/bin && mv /usr/local/bin/migrate.linux-amd64 /usr/local/bin/migrate

WORKDIR /app

COPY go.* ./
RUN go mod download && go mod verify
COPY . .

RUN make build


FROM alpine:latest as server

RUN apk --no-cache add ca-certificates bash

RUN mkdir -p /var/log/app

WORKDIR /app/

COPY --from=build /usr/local/bin/migrate /usr/local/bin
COPY --from=build /app/migrations ./migrations/
COPY --from=build /app/api-server /usr/local/bin
COPY --from=build /app/scripts/entrypoint.sh .

CMD ["/bin/bash", "/app/entrypoint.sh"]
