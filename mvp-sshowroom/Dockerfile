# --- Stage 1: build the Go binary ---
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy the module files first so dependency download is cached separately from
# the source code.
COPY go.mod go.sum ./
RUN go mod download

# Copy the source. ui2/assets/*.yaml must be present here because they are
# embedded into the binary with //go:embed at compile time.
COPY . .

# Build a static binary. "." builds the root main package (the SSH server);
# the ui2 dev runner in ./ui2 is a separate package and is not included.
# -trimpath strips local paths; CGO off keeps the binary portable to Alpine.
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o ssh-server .

# --- Stage 2: minimal runtime image ---
FROM alpine:latest

WORKDIR /app

# Bubble Tea / lipgloss read the terminfo database to match the SSH client's
# TERM; without it the TUI can render with wrong colors or key handling.
RUN apk add --no-cache ncurses-terminfo-base

# Copy the binary from the build stage.
COPY --from=builder /app/ssh-server .

# The SSH server listens on this port (see sshserver/config.go).
EXPOSE 2222

# Start the SSH server.
CMD ["./ssh-server"]
