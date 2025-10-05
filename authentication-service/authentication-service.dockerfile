FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o authApp ./cmd/api
RUN chmod +x /app/authApp

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/authApp /app/
CMD ["/app/authApp"]