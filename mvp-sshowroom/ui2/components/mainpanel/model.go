// Package mainpanel renders the content for the selected topic.
//
// It owns the topic/content state and delegates the actual body to a
// layouts.Layout strategy, so it does not need to know how cards or a CV are
// drawn.
package mainpanel

import (
	tea "github.com/charmbracelet/bubbletea"

	"mvp-sshowroom/ui2/data"
	"mvp-sshowroom/ui2/styles"
)

// Model holds the selected topic and the loaded content used to render it.
type Model struct {
	styles styles.Styles

	topic   data.Topic
	content data.Content
	// hasTopic distinguishes "no topic selected yet" from the zero-value Topic.
	hasTopic bool
}

// New returns an empty main panel bound to the session styles.
func New(st styles.Styles) Model {
	return Model{styles: st}
}

// SetContent stores the loaded dataset so the panel can look up a topic's cards
// or CV entries when it renders.
func (m Model) SetContent(content data.Content) Model {
	m.content = content
	return m
}

// SetTopic switches the topic rendered by the panel.
func (m Model) SetTopic(topic data.Topic) Model {
	m.topic = topic
	m.hasTopic = true
	return m
}

// Init satisfies tea.Model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update satisfies tea.Model. The main panel is currently display-only; all
// navigation happens in the side panel, so it ignores messages.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}
