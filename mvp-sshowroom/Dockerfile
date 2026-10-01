# --- Stage 1: Build dell'eseguibile Go ---
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copia i file dei moduli e scarica le dipendenze
COPY go.mod go.sum ./
RUN go mod download

# Copia il codice sorgente e compila il binario
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ssh-server main.go

# --- Stage 2: Immagine finale minimale ---
FROM alpine:latest

WORKDIR /app

# Copia il binario dallo stage di build
COPY --from=builder /app/ssh-server .

# Espone la porta del server SSH
EXPOSE 2222

# Avvia l'applicazione Go
CMD ["./ssh-server"]
