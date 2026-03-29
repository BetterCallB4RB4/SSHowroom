package ui

import (
	"testing"

	"github.com/alexcloudborn/ssh-portfolio/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func getTestModel() Model {
	r := lipgloss.NewRenderer(nil)
	st := styles.NewStyles(r)
	return InitialModel(st)
}

func TestInitialModel(t *testing.T) {
	m := getTestModel()

	if m.Focus != FocusSidebar {
		t.Errorf("expected initial focus to be FocusSidebar, got %v", m.Focus)
	}

	if len(m.SidebarItems) == 0 {
		t.Error("expected sidebar items to be populated")
	}

	if len(m.Sections) == 0 {
		t.Error("expected sections to be populated")
	}
}

func TestUpdateWindowSize(t *testing.T) {
	m := getTestModel()
	newWidth, newHeight := 100, 50
	
	msg := tea.WindowSizeMsg{Width: newWidth, Height: newHeight}
	updatedModel, _ := m.Update(msg)
	m = updatedModel.(Model)

	if m.Width != newWidth || m.Height != newHeight {
		t.Errorf("expected dimensions %dx%d, got %dx%d", newWidth, newHeight, m.Width, m.Height)
	}
}

func TestUpdateNavigation(t *testing.T) {
	m := getTestModel()
	m.SidebarIdx = 0
	m.Focus = FocusSidebar

	// Press Down (j)
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("j")}
	updatedModel, _ := m.Update(msg)
	m = updatedModel.(Model)

	if m.SidebarIdx != 1 {
		t.Errorf("expected sidebar index to be 1, got %d", m.SidebarIdx)
	}

	// Verify content changes (indirectly by checking if key exists)
	selectedKey := m.SidebarItems[m.SidebarIdx]
	if _, ok := m.Sections[selectedKey]; !ok {
		t.Errorf("selected key %s not found in sections", selectedKey)
	}
}
