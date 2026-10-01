package ui

import "github.com/charmbracelet/lipgloss"

// styles holds the session-bound lipgloss styles used to render the TUI.
// They are built from a *lipgloss.Renderer tied to the client's SSH PTY so
// that colors degrade gracefully for clients with limited terminal
// capabilities.
type styles struct {
	// Generic text styles, reused across panels.
	selected lipgloss.Style
	normal   lipgloss.Style
	info     lipgloss.Style

	// Sidebar-specific styles.
	sidebarBox   lipgloss.Style
	sidebarTitle lipgloss.Style

	// Main-panel-specific styles.
	mainBox     lipgloss.Style
	mainTitle   lipgloss.Style
	tabActive   lipgloss.Style
	tabInactive lipgloss.Style
}

func makeStyles(r *lipgloss.Renderer) styles {
	const (
		accent   = lipgloss.Color("#7D56F4")
		success  = lipgloss.Color("#04B575")
		muted    = lipgloss.Color("#626262")
		fg       = lipgloss.Color("#FAFAFA")
		border   = lipgloss.Color("#444444")
		borderOn = lipgloss.Color("#7D56F4")
	)

	return styles{
		selected: r.NewStyle().Bold(true).Foreground(success),
		normal:   r.NewStyle().Foreground(fg),
		info:     r.NewStyle().Foreground(muted),

		sidebarBox: r.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(border).
			Padding(0, 1),
		sidebarTitle: r.NewStyle().Bold(true).Foreground(accent),

		mainBox: r.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderOn).
			Padding(0, 1),
		mainTitle: r.NewStyle().Bold(true).Foreground(accent).MarginBottom(1),

		tabActive: r.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(accent).
			Padding(0, 1),
		tabInactive: r.NewStyle().
			Foreground(muted).
			Padding(0, 1),
	}
}
