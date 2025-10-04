FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o frontApp ./cmd/web
RUN chmod +x /app/frontApp

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/frontApp /app/
COPY ./cmd/web/templates ./cmd/web/templates
CMD ["/app/frontApp"]