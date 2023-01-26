# build
FROM golang:1.19-bullseye as builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/ohagi-dinner-api

## deploy
FROM litestream/litestream:latest
WORKDIR /app

COPY ./litestream.yml /etc/litestream.yml
COPY --from=builder /app/ohagi-dinner-api ./
COPY ./docker-entrypoint.sh ./
RUN apk add bash

ENTRYPOINT ["bash","docker-entrypoint.sh"]
