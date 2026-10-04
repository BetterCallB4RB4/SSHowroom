package sshserver

import (
	"errors"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	wishbubbletea "github.com/charmbracelet/wish/bubbletea"
	wishlogging "github.com/charmbracelet/wish/logging"

	tea "github.com/charmbracelet/bubbletea"

	"mvp-sshowroom/ui2/app"
)

// New builds a Wish SSH server wired up to serve the SSHowroom TUI, using an
// in-memory generated host key and the given configuration.
func New(cfg Config) (*ssh.Server, error) {
	hostKeyPEM, err := generateInMemoryHostKey()
	if err != nil {
		return nil, err
	}

	return wish.NewServer(
		wish.WithAddress(cfg.Addr()),
		wish.WithHostKeyPEM(hostKeyPEM),
		wish.WithMiddleware(
			wishbubbletea.Middleware(teaHandler),
			wishlogging.Middleware(),
		),
	)
}

// teaHandler bridges an incoming SSH session to a Bubble Tea program backed
// by the SSHowroom TUI model.
func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	_, _, active := s.Pty()
	if !active {
		wish.Fatalln(s, errors.New("no active terminal PTY found"))
		return nil, nil
	}

	// Make renderer derived specifically from this client's SSH PTY session
	renderer := wishbubbletea.MakeRenderer(s)
	m := app.New(renderer)

	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
