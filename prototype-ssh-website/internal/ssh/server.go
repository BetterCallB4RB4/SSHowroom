package ssh

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alexcloudborn/ssh-portfolio/internal/styles"
	"github.com/alexcloudborn/ssh-portfolio/internal/ui"
	"github.com/muesli/termenv"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

func NewSSHServer(host string, port int) (*ssh.Server, error) {
	return wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				pty, _, active := s.Pty()
				if !active {
					return nil, nil
				}

				// Renderer per sessione
				renderer := lipgloss.NewRenderer(s)

				// Forziamo il profilo colore usando termenv
				term := strings.ToLower(pty.Term)
				if strings.Contains(term, "truecolor") || strings.Contains(term, "xterm") || strings.Contains(term, "iterm") {
					renderer.SetColorProfile(termenv.TrueColor)
				} else if strings.Contains(term, "256color") {
					renderer.SetColorProfile(termenv.ANSI256)
				} else {
					renderer.SetColorProfile(termenv.ANSI)
				}

				st := styles.NewStyles(renderer)
				return ui.InitialModel(st), []tea.ProgramOption{tea.WithAltScreen()}
			}),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
}

func Shutdown(s *ssh.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return s.Shutdown(ctx)
}
