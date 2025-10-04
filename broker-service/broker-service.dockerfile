FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api
RUN chmod +x /app/brokerApp

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/brokerApp /app/
CMD ["/app/brokerApp"]