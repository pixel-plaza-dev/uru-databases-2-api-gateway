FROM golang:1.23.2-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency --parseInternal
RUN go build -v -o /bin/api .

FROM alpine:latest AS api

WORKDIR /app

COPY --from=builder /bin/api /api

EXPOSE 8080

ENTRYPOINT ["/api", "-m=prod"]