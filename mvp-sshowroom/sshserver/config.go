package sshserver

import (
	"fmt"
	"net"
)

// Config holds the settings needed to start the SSH server.
type Config struct {
	Host string
	Port int
}

// DefaultConfig returns the default server configuration.
func DefaultConfig() Config {
	return Config{
		Host: "0.0.0.0",
		Port: 2222,
	}
}

// Addr returns the "host:port" address derived from the configuration.
func (c Config) Addr() string {
	return net.JoinHostPort(c.Host, fmt.Sprintf("%d", c.Port))
}
