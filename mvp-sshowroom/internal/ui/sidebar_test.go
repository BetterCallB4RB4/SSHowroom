package ui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestSidebarTabCyclesTopics(t *testing.T) {
	s := newSidebar([]Topic{{Name: "Projects"}, {Name: "About"}, {Name: "Contact"}})
	s = s.Move(1)
	if got := s.Selected().Name; got != "About" {
		t.Fatalf("Tab selected %q, want About", got)
	}
	s = s.Move(-1)
	if got := s.Selected().Name; got != "Projects" {
		t.Fatalf("Shift+Tab selected %q, want Projects", got)
	}
}

func TestSidebarArrowNavigation(t *testing.T) {
	s := newSidebar([]Topic{{Name: "Projects"}, {Name: "About"}})
	updated, changed := s.Update(tea.KeyMsg{Type: tea.KeyDown})
	if !changed || updated.Selected().Name != "About" {
		t.Fatalf("down did not select About: changed=%v selected=%q", changed, updated.Selected().Name)
	}
}
