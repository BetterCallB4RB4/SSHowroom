package ui

import (
	"math/rand"

	"github.com/charmbracelet/lipgloss"
)

var footerColors = []lipgloss.Color{
	lipgloss.Color("#FF5F5F"),
	lipgloss.Color("#FFD75F"),
	lipgloss.Color("#5FFF87"),
	lipgloss.Color("#5FD7FF"),
	lipgloss.Color("#AF87FF"),
	lipgloss.Color("#FF87D7"),
	lipgloss.Color("#FFFFFF"),
}

func nextFooterColor(current lipgloss.Color) lipgloss.Color {
	next := rand.Intn(len(footerColors))
	if footerColors[next] == current {
		next = (next + 1 + rand.Intn(len(footerColors)-1)) % len(footerColors)
	}
	return footerColors[next]
}

// styles holds the session-bound lipgloss styles used to render the TUI.
// They are built from a *lipgloss.Renderer tied to the client's SSH PTY so
// that colors degrade gracefully for clients with limited terminal
// capabilities.
type styles struct {
	// Generic text styles, reused across panels.
	selected lipgloss.Style
	normal   lipgloss.Style
	info     lipgloss.Style

	// Shared panel styles keep the dashboard visually consistent.
	panel           lipgloss.Style
	panelTitle      lipgloss.Style
	card            lipgloss.Style
	selectedCard    lipgloss.Style
	cardTitle       lipgloss.Style
	cardUnderline   lipgloss.Style
	cardDescription lipgloss.Style
	cardLink        lipgloss.Style

	// Main-panel-specific styles.
	tabActive   lipgloss.Style
	tabInactive lipgloss.Style

	// Animated face panel styles.
	faceWhite  lipgloss.Style
	faceIris   lipgloss.Style
	facePupil  lipgloss.Style
	faceLid    lipgloss.Style
	faceMouth  lipgloss.Style
	faceAccent lipgloss.Color
	faceHappy  lipgloss.Color

	// Ghosttime animation styles.
	ghostOutline lipgloss.Style
	ghostBody    lipgloss.Style
	ghostColors  []lipgloss.Color
}

func makeStyles(r *lipgloss.Renderer) styles {
	const (
		accent    = lipgloss.Color("#7D56F4")
		selection = lipgloss.Color("#FFD75F")
		black     = lipgloss.Color("#000000")
		white     = lipgloss.Color("#FFFFFF")
		muted     = lipgloss.Color("#626262")
		fg        = lipgloss.Color("#FAFAFA")
		faceBlue  = lipgloss.Color("#67D9E8")
		faceGold  = lipgloss.Color("#FFD75F")
	)

	return styles{
		selected: r.NewStyle().
			Bold(true).
			Foreground(black).
			Background(selection).
			Padding(0, 2),
		normal: r.NewStyle().Foreground(fg),
		info:   r.NewStyle().Foreground(muted),

		panel: r.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(white).
			Padding(0, 1),
		panelTitle: r.NewStyle().Bold(true).Foreground(accent),
		card: r.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(muted).
			Padding(0, 1),
		selectedCard: r.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(selection).
			Padding(0, 1),
		cardTitle:       r.NewStyle().Bold(true).Foreground(fg),
		cardUnderline:   r.NewStyle().Bold(true),
		cardDescription: r.NewStyle().Foreground(fg),
		cardLink:        r.NewStyle().Foreground(selection).Underline(true),

		tabActive: r.NewStyle().
			Bold(true).
			Foreground(black).
			Background(selection),
		tabInactive: r.NewStyle().
			Foreground(muted),

		faceWhite:  r.NewStyle().Background(lipgloss.Color("#DDFBFF")),
		faceIris:   r.NewStyle().Background(lipgloss.Color("#27B6C7")),
		facePupil:  r.NewStyle().Background(lipgloss.Color("#092B3A")),
		faceLid:    r.NewStyle().Background(lipgloss.Color("#7D56F4")),
		faceMouth:  r.NewStyle().Bold(true),
		faceAccent: faceBlue,
		faceHappy:  faceGold,

		ghostOutline: r.NewStyle().Bold(true),
		ghostBody: r.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#D8D6E0")),
		ghostColors: []lipgloss.Color{
			lipgloss.Color("#79A8FF"),
			lipgloss.Color("#B38CFF"),
			lipgloss.Color("#FF78B7"),
			lipgloss.Color("#FF8B65"),
			lipgloss.Color("#78E0A1"),
			lipgloss.Color("#64D8E8"),
		},
	}
}
