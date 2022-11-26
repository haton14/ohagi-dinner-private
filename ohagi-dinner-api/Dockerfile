FROM golang:1.19 as builder

WORKDIR /app
COPY ./ /app
RUN go mod download -x
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/ohagi-dinner-api

FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/ohagi-dinner-api /
ENV PORT=${PORT}
ENTRYPOINT ["/ohagi-dinner-api"]
