# build
FROM golang:1.19-bullseye as builder
WORKDIR /app
COPY ./ /app
RUN go mod download -x
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/ohagi-dinner-api

## deploy
FROM litestream/litestream:latest
WORKDIR /app
COPY my-secret.json /app/.gcs-credentials.json
RUN cat /app/.gcs-credentials.json
ENV GOOGLE_APPLICATION_CREDENTIALS /app/.gcs-credentials.json
ENV TZ Asia/Tokyo

COPY ./litestream.yml /etc/litestream.yml
COPY --from=builder /app/ohagi-dinner-api /app
COPY ./docker-entrypoint.sh /app
RUN apk add bash

ENTRYPOINT ["bash","docker-entrypoint.sh"]
