package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func TestMainPanelGridNavigation(t *testing.T) {
	topic := Topic{Cards: make([]Card, 5)}
	m := newMainPanel()

	m = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'l'}}, topic, 2)
	if m.selectedCard != 1 {
		t.Fatalf("right should select card 1, got %d", m.selectedCard)
	}
	m = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}}, topic, 2)
	if m.selectedCard != 3 {
		t.Fatalf("down should select card 3, got %d", m.selectedCard)
	}
	m = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}}, topic, 2)
	if m.selectedCard != 3 {
		t.Fatalf("down at final row should clamp, got %d", m.selectedCard)
	}
	m = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'h'}}, topic, 2)
	if m.selectedCard != 2 {
		t.Fatalf("left should select card 2, got %d", m.selectedCard)
	}
}

func TestMainPanelEnterShowsSelectedURL(t *testing.T) {
	topic := Topic{Cards: []Card{{URL: "https://google.com"}}}
	m := newMainPanel().Update(tea.KeyMsg{Type: tea.KeyEnter}, topic, 1)
	if m.openURL != "https://google.com" {
		t.Fatalf("enter selected %q, want URL", m.openURL)
	}
}

func TestMainPanelRendersCardContentAndHyperlink(t *testing.T) {
	r := lipgloss.NewRenderer(nil)
	st := makeStyles(r)
	topic := Topic{Name: "Projects", Cards: []Card{{
		Title:       "Showroom",
		Description: "A portfolio in the terminal",
		URL:         "https://google.com",
		Underline:   lipgloss.Color("#FFD75F"),
	}}}
	view := newMainPanel().View(st, 50, 20, topic)
	for _, text := range []string{"Projects", "Showroom", "A portfolio in the terminal", "https://google.com", "\x1b]8;;https://google.com"} {
		if !strings.Contains(view, text) {
			t.Errorf("rendered card is missing %q", text)
		}
	}
}
