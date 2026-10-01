package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/ssh"

	"mvp-sshowroom/internal/sshserver"
)

func main() {
	cfg := sshserver.DefaultConfig()

	server, err := sshserver.New(cfg)
	if err != nil {
		log.Fatalln("Failed to create Wish server:", err)
	}

	// Allows the application to shut down cleanly from the system or Kubernetes.
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("🚀 SSH Server listening on %s...", cfg.Addr())
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Fatalln("Server error:", err)
		}
	}()

	<-done
	log.Println("Shutting down SSH server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("Shutdown error:", err)
	}
}
