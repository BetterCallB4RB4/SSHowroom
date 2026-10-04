// Package bigfont turns a short string into multi-row ASCII art so headings
// look larger than the rest of the UI.
//
// Terminals pick the glyph size, so a TUI cannot actually change the font.
// Drawing each letter several rows tall is the closest a server-side app can
// get to "a bigger font".
package bigfont

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
)

// fonts is tried in order from widest/nicest to most compact. The first one
// whose art fits the available width wins.
var fonts = []string{"standard", "small", "mini"}

// Render returns text as ASCII art that fits within maxWidth columns.
//
// It tries progressively narrower figlet fonts and, if even the smallest does
// not fit, falls back to the plain string so a narrow terminal still shows a
// readable title instead of a wrapped mess.
func Render(text string, maxWidth int) string {
	for _, font := range fonts {
		art := strings.Trim(figure.NewFigure(text, font, true).String(), "\n")
		if art != "" && Widest(art) <= maxWidth {
			return art
		}
	}
	return text
}

// Widest returns the display width of the longest line in s.
func Widest(s string) int {
	widest := 0
	for _, line := range strings.Split(s, "\n") {
		if w := lipgloss.Width(line); w > widest {
			widest = w
		}
	}
	return widest
}
