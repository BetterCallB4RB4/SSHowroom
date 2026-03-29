# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
# ENV TERM=xterm-256color
# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o portfolio main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Install openssh-keygen and ncurses-terminfo-base for terminal support
RUN apk add --no-cache openssh-keygen ncurses-terminfo-base

# Create the .ssh directory and set permissions
RUN mkdir -p .ssh && chmod 700 .ssh

# Copy the binary from the build stage
COPY --from=builder /app/portfolio .

# The application runs on port 2222
EXPOSE 2222

# Run the application
CMD ["./portfolio"]
