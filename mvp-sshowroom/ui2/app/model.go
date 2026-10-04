// Package app wires the panels together into the root Bubble Tea model,
// following the Elm Architecture: a Model holding all state, an Init that
// returns startup commands, an Update that folds messages into new state, and a
// View that renders state to a string.
package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"mvp-sshowroom/ui2/components/companionpanel"
	"mvp-sshowroom/ui2/components/mainpanel"
	"mvp-sshowroom/ui2/components/sidepanel"
	"mvp-sshowroom/ui2/data"
	"mvp-sshowroom/ui2/styles"
)

// Model is the root model. It owns the terminal dimensions and the three panel
// sub-models, and it routes messages between them.
type Model struct {
	width  int
	height int

	side      sidepanel.Model
	main      mainpanel.Model
	companion companionpanel.Model

	// err is non-nil when the content files failed to load; View surfaces it.
	err error
}

// New builds the root model for one SSH session. The renderer is derived from
// the client's PTY so colors degrade gracefully per client.
func New(r *lipgloss.Renderer) Model {
	st := styles.New(r)
	return Model{
		side:      sidepanel.New(st),
		main:      mainpanel.New(st),
		companion: companionpanel.New(st),
	}
}

// Init kicks off the asynchronous content load.
//
// The YAML is embedded in the binary, so this finishes almost instantly, but
// wrapping it in a tea.Cmd keeps all I/O off the update path and matches the
// Elm pattern (Init returns commands, Update handles their messages).
func (m Model) Init() tea.Cmd {
	return loadContentCmd()
}

// dataLoadedMsg carries the result of a successful load.
type dataLoadedMsg struct {
	content data.Content
}

// dataLoadFailedMsg carries a load error.
type dataLoadFailedMsg struct {
	err error
}

// loadContentCmd returns a command that reads and decodes the embedded content.
// The actual work happens when Bubble Tea runs the returned function.
func loadContentCmd() tea.Cmd {
	return func() tea.Msg {
		content, err := data.Load()
		if err != nil {
			return dataLoadFailedMsg{err: err}
		}
		return dataLoadedMsg{content: content}
	}
}
