package styles

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Renderer *lipgloss.Renderer

	// Colors
	TsBlack  lipgloss.TerminalColor
	TsWhite  lipgloss.TerminalColor
	TsOrange lipgloss.TerminalColor
	TsGray   lipgloss.TerminalColor

	// Layout Containers
	HeaderStyle  lipgloss.Style
	FooterStyle  lipgloss.Style
	SidebarStyle lipgloss.Style
	MainBoxStyle lipgloss.Style

	// UI Elements
	HeaderItemStyle   lipgloss.Style
	HeaderActiveStyle lipgloss.Style
	
	SidebarItemStyle     lipgloss.Style
	SidebarSelectedStyle lipgloss.Style
	
	ContentTitleStyle lipgloss.Style
	ContentBodyStyle  lipgloss.Style
	
	ShortcutKeyStyle  lipgloss.Style
	ShortcutDescStyle lipgloss.Style
}

func NewStyles(r *lipgloss.Renderer) Styles {
	// Colors - Terminal.shop Inspired
	tsBlack := lipgloss.Color("#000000")
	tsWhite := lipgloss.Color("#ffffff")
	tsOrange := lipgloss.Color("#ff5f00") // The "Plus One" accent
	tsGray := lipgloss.Color("#333333")   // Subtle borders

	s := Styles{
		Renderer: r,
		TsBlack:  tsBlack,
		TsWhite:  tsWhite,
		TsOrange: tsOrange,
		TsGray:   tsGray,
	}

	// Thin borders, no rounded corners (Lazygit Functional Structure)
	border := lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
	}

	// --- Layout Containers ---

	s.HeaderStyle = r.NewStyle().
		MarginBottom(1).
		Padding(0, 1)

	s.FooterStyle = r.NewStyle().
		MarginTop(1).
		Padding(0, 1).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		BorderForeground(tsGray)

	s.SidebarStyle = r.NewStyle().
		Border(border).
		BorderForeground(tsGray).
		Padding(1, 1).
		MarginRight(1)

	s.MainBoxStyle = r.NewStyle().
		Border(border).
		BorderForeground(tsGray).
		Padding(1, 2)

	// --- UI Elements ---

	s.HeaderItemStyle = r.NewStyle().
		Foreground(tsWhite).
		MarginRight(2)

	s.HeaderActiveStyle = s.HeaderItemStyle.Copy().
		Foreground(tsOrange).
		Bold(true).
		Underline(true)

	s.SidebarItemStyle = r.NewStyle().
		Foreground(tsWhite).
		PaddingLeft(1).
		MarginBottom(0)

	// Selected blocks (Terminal.shop style)
	s.SidebarSelectedStyle = r.NewStyle().
		Foreground(tsBlack).
		Background(tsOrange).
		Bold(true).
		PaddingLeft(1)

	s.ContentTitleStyle = r.NewStyle().
		Foreground(tsOrange).
		Bold(true).
		MarginBottom(1).
		SetString("● ") // Minimalist bullet

	s.ContentBodyStyle = r.NewStyle().
		Foreground(tsWhite).
		PaddingTop(1)

	s.ShortcutKeyStyle = r.NewStyle().
		Foreground(tsOrange).
		Bold(true)

	s.ShortcutDescStyle = r.NewStyle().
		Foreground(tsWhite)

	return s
}
