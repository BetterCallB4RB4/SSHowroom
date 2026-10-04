// Package styles centralises every lipgloss style used by the TUI.
//
// Styles are built from a *lipgloss.Renderer that is tied to a single SSH
// client's PTY, so colors degrade gracefully for clients with limited terminal
// capabilities. Because lipgloss styles are immutable values, callers derive
// per-topic variants (for example a colored title) with Foreground(...) without
// mutating the shared base style.
package styles

import (
	"math"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	// DefaultAccent is used when a topic does not specify a color.
	DefaultAccent = lipgloss.Color("#7D56F4")
	// white is the border color required by the layout spec.
	white = lipgloss.Color("#FFFFFF")
	// muted is used for secondary text such as dates.
	muted = lipgloss.Color("#8A8A8A")
	// shadow is the color of the card drop shadow.
	shadow = lipgloss.Color("#3A3A3A")
)

// Panel and card padding. These are exported because the panels need them to
// know how much of their content box is usable for text (lipgloss counts
// padding as part of Width/Height).
const (
	PanelPaddingX = 2
	PanelPaddingY = 1
	CardPaddingX  = 2
)

// shadowBorder draws only the right and bottom edges of a card using half-block
// characters. Painted in a dim color it reads as a drop shadow, so the card
// body itself stays transparent instead of being filled.
var shadowBorder = lipgloss.Border{
	Right:       "▐",
	Bottom:      "▄",
	BottomRight: "▟",
}

// Styles holds the session-bound base styles. Panels copy these and recolor
// them with the active topic's color.
type Styles struct {
	// Panel is the outer box shared by all three panels: a thick, square,
	// white border with no background (transparent fill) and inner padding so
	// text never touches the border.
	Panel lipgloss.Style
	// Card is the transparent card body with a soft drop shadow on its right
	// and bottom edges.
	Card lipgloss.Style
	// Title is used for panel and card titles. Callers add Foreground+Underline
	// to get a colored underline (lipgloss underlines inherit the text color).
	Title lipgloss.Style
	// Item is normal body text.
	Item lipgloss.Style
	// Selected marks the highlighted row in the side panel.
	Selected lipgloss.Style
	// Muted is secondary text.
	Muted lipgloss.Style
	// Link styles external links.
	Link lipgloss.Style
}

// New builds the base Styles from a client's renderer.
func New(r *lipgloss.Renderer) Styles {
	return Styles{
		// No Background(...) call keeps the panels transparent.
		Panel: r.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(white).
			Padding(PanelPaddingY, PanelPaddingX),
		Card: r.NewStyle().
			Padding(1, CardPaddingX).
			Border(shadowBorder, false, true, true, false).
			BorderForeground(shadow),
		Title:    r.NewStyle().Bold(true),
		Item:     r.NewStyle(),
		Selected: r.NewStyle().Bold(true),
		Muted:    r.NewStyle().Foreground(muted),
		Link:     r.NewStyle().Underline(true),
	}
}

// Accent converts a topic's YAML color string into a lipgloss color, falling
// back to DefaultAccent when the value is empty.
func Accent(hex string) lipgloss.Color {
	if hex == "" {
		return DefaultAccent
	}
	return lipgloss.Color(hex)
}

// OnAccent picks a foreground color (black or white) that stays readable when
// drawn on top of the given background color. It compares the WCAG contrast
// ratio against white and against black and returns whichever is higher, so a
// bright accent gets black text and a dark accent gets white text.
func OnAccent(bg lipgloss.Color) lipgloss.Color {
	lum := relativeLuminance(string(bg))

	// +0.05 is the WCAG offset; 1.05 is the contrast of white (luminance 1).
	contrastWithWhite := 1.05 / (lum + 0.05)
	contrastWithBlack := (lum + 0.05) / 0.05

	if contrastWithBlack >= contrastWithWhite {
		return lipgloss.Color("#000000")
	}
	return lipgloss.Color("#FFFFFF")
}

// relativeLuminance returns the WCAG relative luminance (0 = black, 1 = white)
// of a #RGB or #RRGGBB color. Unparseable values are treated as white so the
// caller falls back to black text.
func relativeLuminance(hex string) float64 {
	r, g, b := parseHex(hex)
	return 0.2126*linearize(r) + 0.7152*linearize(g) + 0.0722*linearize(b)
}

// parseHex converts a hex color into normalized (0..1) red, green and blue.
func parseHex(hex string) (float64, float64, float64) {
	hex = strings.TrimPrefix(strings.TrimSpace(hex), "#")
	// Expand the shorthand #abc form into #aabbcc.
	if len(hex) == 3 {
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}
	if len(hex) != 6 {
		return 1, 1, 1
	}

	value, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return 1, 1, 1
	}
	return float64(value>>16&0xff) / 255,
		float64(value>>8&0xff) / 255,
		float64(value&0xff) / 255
}

// linearize converts one sRGB channel to its linear-light value, as required
// by the WCAG luminance formula.
func linearize(channel float64) float64 {
	if channel <= 0.03928 {
		return channel / 12.92
	}
	return math.Pow((channel+0.055)/1.055, 2.4)
}
