package sshserver

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"

	gossh "golang.org/x/crypto/ssh"
)

// generateInMemoryHostKey creates a fresh ed25519 SSH host key, PEM-encoded,
// kept only in memory (no key is written to disk).
func generateInMemoryHostKey() ([]byte, error) {
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	pemBlock, err := gossh.MarshalPrivateKey(privateKey, "")
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(pemBlock), nil
}
