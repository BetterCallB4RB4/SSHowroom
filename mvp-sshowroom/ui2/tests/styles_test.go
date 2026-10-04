package tests

import (
	"testing"

	"github.com/charmbracelet/lipgloss"

	"mvp-sshowroom/ui2/styles"
)

// TestOnAccent checks the readable-foreground picker. Black and white are
// unambiguous extremes; the purple is the app's default accent, which is dark
// enough that white text reads better than black.
func TestOnAccent(t *testing.T) {
	cases := []struct {
		bg   lipgloss.Color
		want lipgloss.Color
	}{
		{"#000000", lipgloss.Color("#FFFFFF")}, // dark background
		{"#FFFFFF", lipgloss.Color("#000000")}, // light background
		{"#7D56F4", lipgloss.Color("#FFFFFF")}, // default accent
		{"#04B575", lipgloss.Color("#000000")}, // bright green
	}

	for _, tc := range cases {
		if got := styles.OnAccent(tc.bg); got != tc.want {
			t.Errorf("OnAccent(%s) = %s, want %s", tc.bg, got, tc.want)
		}
	}
}
