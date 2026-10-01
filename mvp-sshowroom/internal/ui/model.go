package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model is the root Bubble Tea model backing the SSHowroom TUI. It follows
// the Elm architecture by composing two smaller sub-models — the sidebar
// and the main panel — and coordinating messages between them.
type Model struct {
	topics []Topic

	sidebar sidebarModel
	main    mainModel

	width  int
	height int
	ready  bool // true once the first tea.WindowSizeMsg has arrived

	quitting bool
	styles   styles
}

// NewModel builds the initial Model using the renderer derived from the
// client's SSH PTY session. The renderer is required so that Lipgloss
// styles respect each individual client's color profile/capabilities.
func NewModel(r *lipgloss.Renderer) Model {
	topics := mockTopics()

	return Model{
		topics:  topics,
		sidebar: newSidebar(topics),
		main:    newMainPanel(),
		styles:  makeStyles(r),
	}
}

// Init satisfies tea.Model. No startup command is needed here, but it's
// kept as a dedicated method since real apps often kick off tea.Cmds (e.g.
// loading data asynchronously) from here.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update is the single entry point for all state transitions. It dispatches
// incoming messages to the relevant sub-model(s) and returns the updated
// Model plus any tea.Cmd that should be executed next.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// tea.WindowSizeMsg arrives whenever the terminal is resized (including
	// the very first time the program starts). We store the dimensions so
	// View() can lay out the two panels proportionally.
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}

		// Sidebar navigation (j/k) — selecting a new topic resets the main
		// panel's active tab, since tab indices are only meaningful per
		// topic (e.g. "Contact" has fewer tabs than "Projects").
		var topicChanged bool
		m.sidebar, topicChanged = m.sidebar.Update(msg)
		if topicChanged {
			m.main = m.main.ResetTab()
			return m, nil
		}

		// Main panel tab navigation (h/l) — scoped to whichever topic is
		// currently selected in the sidebar.
		m.main = m.main.Update(msg, m.sidebar.Selected())
	}

	return m, nil
}
