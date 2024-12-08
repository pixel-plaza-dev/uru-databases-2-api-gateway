FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /bin/api .

FROM alpine:latest AS api

WORKDIR /app

COPY --from=builder /bin/api /api

EXPOSE 8080

ENTRYPOINT ["/api", "-m=prod"]