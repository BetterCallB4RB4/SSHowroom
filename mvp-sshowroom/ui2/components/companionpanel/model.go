// Package companionpanel renders the top-right panel of the layout.
//
// It is intentionally empty for now: it exists only so the grid has a third
// cell, and it is the natural home for future content (status, clock, etc.).
package companionpanel

import (
	tea "github.com/charmbracelet/bubbletea"

	"mvp-sshowroom/ui2/styles"
)

// Model is a placeholder panel. It stores the session styles so its View can
// draw the same border as the other panels.
type Model struct {
	styles styles.Styles
}

// New returns an empty companion panel bound to the session styles.
func New(st styles.Styles) Model {
	return Model{styles: st}
}

// Init satisfies tea.Model. The panel has no startup work.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update satisfies tea.Model. The panel currently ignores every message.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}
