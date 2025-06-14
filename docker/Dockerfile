FROM golang:1.21-alpine AS builder

ENV CGO_ENABLED=0 \
    GO111MODULE=on

WORKDIR /app

RUN apk update && apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN go build -o websocket-server ./cmd/websocket-server



# Create working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/websocket-server .

# Expose port for WebSocket server
EXPOSE 8084

# Run the WebSocket server
ENTRYPOINT ["./websocket-server"]