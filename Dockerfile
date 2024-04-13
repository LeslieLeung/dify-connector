FROM golang:1.22.2-alpine3.19 AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o dify-connector .

FROM alpine:3.19.1

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

USER nonroot

WORKDIR /app

COPY --from=builder /app/dify-connector /app/dify-connector

CMD ["/app/dify-connector", "serve"]