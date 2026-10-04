package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"mvp-sshowroom/ui2/components/sidepanel"
)

// Update is the single funnel for every message, per the Elm Architecture.
//
// It handles the messages the root owns (resize, quit, data loading) and
// forwards key presses to the side panel, which is the only interactive
// component. When the side panel reports a new selection, the root updates the
// main panel -- this is how the topic-choice flows from one component to
// another without the components importing each other.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		// Remember the PTY size; the next View recomputes the grid from it.
		m.width, m.height = msg.Width, msg.Height
		return m, nil

	case dataLoadedMsg:
		m.err = nil
		m.side = m.side.SetTopics(msg.content.Topics)
		m.main = m.main.SetContent(msg.content)
		// Select the first topic so the main panel is never empty, and so the
		// side panel's highlighted row matches what is displayed.
		if len(msg.content.Topics) > 0 {
			m.main = m.main.SetTopic(msg.content.Topics[0])
		}
		return m, nil

	case dataLoadFailedMsg:
		m.err = msg.err
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		// Only the side panel reacts to navigation keys; it returns a command
		// carrying TopicSelectedMsg if the cursor moved.
		var cmd tea.Cmd
		m.side, cmd = m.side.Update(msg)
		return m, cmd

	case sidepanel.TopicSelectedMsg:
		m.main = m.main.SetTopic(msg.Topic)
		return m, nil
	}

	return m, nil
}
