package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model is the root Bubble Tea model backing the SSHowroom TUI. It follows
// the Elm architecture by composing the sidebar, main panel, animated eyes,
// and spinning ghost panel, and coordinating messages between them.
type Model struct {
	topics []Topic

	sidebar  sidebarModel
	main     mainModel
	face     faceModel
	terminal terminalModel

	splash      bool
	width       int
	height      int
	ready       bool // true once the first tea.WindowSizeMsg has arrived
	footerColor lipgloss.Color

	quitting bool
	styles   styles
}

// NewModel builds the initial Model using the renderer derived from the
// client's SSH PTY session. The renderer is required so that Lipgloss
// styles respect each individual client's color profile/capabilities.
func NewModel(r *lipgloss.Renderer) Model {
	topics := mockTopics()

	return Model{
		topics:      topics,
		sidebar:     newSidebar(topics),
		main:        newMainPanel(),
		face:        newFace(),
		splash:      true,
		footerColor: nextFooterColor(""),
		styles:      makeStyles(r),
	}
}

// Init satisfies tea.Model. No startup command is needed here, but it's
// kept as a dedicated method since real apps often kick off tea.Cmds (e.g.
// loading data asynchronously) from here.
func (m Model) Init() tea.Cmd {
	return nil
}

type faceTickMsg time.Time

func nextFaceTick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return faceTickMsg(t)
	})
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

	case faceTickMsg:
		if m.splash || m.quitting {
			return m, nil
		}
		m.face = m.face.Update(time.Time(msg))
		m.terminal = m.terminal.Update()
		return m, nextFaceTick()

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			m.quitting = true
			return m, tea.Quit
		}
		if m.splash {
			m.splash = false
			return m, nextFaceTick()
		}

		m.footerColor = nextFooterColor(m.footerColor)
		m.face = m.face.LookAt(msg.String(), time.Now())
		if msg.String() == "q" {
			m.quitting = true
			return m, tea.Quit
		}
		if msg.String() == "tab" || msg.String() == "shift+tab" {
			delta := 1
			if msg.String() == "shift+tab" {
				delta = -1
			}
			m.sidebar = m.sidebar.Move(delta)
			m.main = m.main.ResetSelection()
			return m, nil
		}

		cardWidth := 34
		if m.width >= 115 {
			cardWidth = max(1, (m.width-int(float64(m.width)*sidebarWidthRatio)-panelGap-6)/2)
		}
		columns := max(1, (m.width-int(float64(m.width)*sidebarWidthRatio)-panelGap-4)/cardWidth)
		m.main = m.main.Update(msg, m.sidebar.Selected(), columns)
	}

	return m, nil
}
