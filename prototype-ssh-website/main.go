package main

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexcloudborn/ssh-portfolio/internal/ssh"
	"github.com/charmbracelet/log"
	sshlib "github.com/charmbracelet/ssh"
)

const (
	host = "0.0.0.0"
	port = 2222
)

func main() {
	s, err := ssh.NewSSHServer(host, port)
	if err != nil {
		log.Fatal("Could not start server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, sshlib.ErrServerClosed) {
			log.Error("Could not stop server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	if err := ssh.Shutdown(s); err != nil {
		log.Error("Error during shutdown", "error", err)
	}
}
