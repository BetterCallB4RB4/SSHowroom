// Package sidepanel renders the navigable list of topics.
package sidepanel

import (
	tea "github.com/charmbracelet/bubbletea"

	"mvp-sshowroom/ui2/data"
	"mvp-sshowroom/ui2/styles"
)

// Model holds the topic list and the cursor position.
type Model struct {
	styles styles.Styles
	topics []data.Topic
	cursor int
}

// New returns an empty side panel bound to the session styles.
func New(st styles.Styles) Model {
	return Model{styles: st}
}

// SetTopics replaces the topic list. The cursor is clamped to the new length so
// reusing the panel with a shorter list cannot leave it pointing out of range.
func (m Model) SetTopics(topics []data.Topic) Model {
	m.topics = topics
	if m.cursor >= len(topics) {
		m.cursor = 0
	}
	return m
}

// Selected returns the currently highlighted topic.
func (m Model) Selected() (data.Topic, bool) {
	if m.cursor < 0 || m.cursor >= len(m.topics) {
		return data.Topic{}, false
	}
	return m.topics[m.cursor], true
}

// Init satisfies tea.Model.
func (m Model) Init() tea.Cmd {
	return nil
}

// TopicSelectedMsg is emitted whenever the cursor moves, so the root model can
// tell the main panel which topic to render.
//
// Carrying the whole Topic (rather than just an index) keeps the side panel and
// the main panel decoupled: the main panel never needs to know about the list.
type TopicSelectedMsg struct {
	Topic data.Topic
}

// Update handles navigation. Tab moves forward and Shift+Tab moves back; both
// wrap around. When the cursor moves it returns a command carrying a
// TopicSelectedMsg, which the root model routes to the main panel.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok || len(m.topics) == 0 {
		return m, nil
	}

	switch key.String() {
	case "tab":
		m.cursor = (m.cursor + 1) % len(m.topics)
	case "shift+tab":
		m.cursor = (m.cursor - 1 + len(m.topics)) % len(m.topics)
	default:
		return m, nil
	}

	topic := m.topics[m.cursor]
	return m, func() tea.Msg { return TopicSelectedMsg{Topic: topic} }
}
