# ===== BUILD STAGE =====
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git, bash, ca-certificates
RUN apk add --no-cache git bash ca-certificates

# Set certs environment
ENV GIT_SSL_NO_VERIFY=0

# Copy dependencies & download
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app


# ===== RUNTIME STAGE =====
FROM alpine:3.19

WORKDIR /app

# Certificates
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/app .

COPY .env ./

# Expose port API
EXPOSE 5000

# Run the built binary
ENTRYPOINT ["./app"]
CMD ["serve"]
