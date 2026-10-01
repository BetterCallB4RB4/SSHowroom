package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// mainModel is the state for the right-hand content panel. It tracks which
// Activity "tab" is active for the currently selected Topic. It knows
// nothing about the sidebar beyond the Topic it is handed to render.
type mainModel struct {
	activeTab int
}

// newMainPanel builds the initial main-panel state.
func newMainPanel() mainModel {
	return mainModel{}
}

// ResetTab snaps the active tab back to the first one. The root Model calls
// this whenever the sidebar selection changes, since tab index 3 on
// "Projects" may not exist on "Contact".
func (m mainModel) ResetTab() mainModel {
	m.activeTab = 0
	return m
}

// Update handles tab navigation within the main panel. Unlike the sidebar,
// movement here is clamped rather than wrapped: trying to go past the last
// tab simply stays on the last tab, which better matches how tab bars
// typically behave.
func (m mainModel) Update(msg tea.Msg, topic Topic) mainModel {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m
	}

	switch keyMsg.String() {
	case "h", "left":
		if m.activeTab > 0 {
			m.activeTab--
		}

	case "l", "right":
		if m.activeTab < len(topic.Activities)-1 {
			m.activeTab++
		}
	}

	return m
}

// View renders the tab bar followed by the active Activity's content, boxed
// to the given width/height.
func (m mainModel) View(st styles, width, height int, topic Topic) string {
	var tabs []string
	for i, activity := range topic.Activities {
		if i == m.activeTab {
			tabs = append(tabs, st.tabActive.Render(activity.Title))
		} else {
			tabs = append(tabs, st.tabInactive.Render(activity.Title))
		}
	}
	tabBar := strings.Join(tabs, " ")

	var content string
	if len(topic.Activities) > 0 {
		content = topic.Activities[m.activeTab].Content
	}

	body := st.mainTitle.Render(topic.Name) + "\n" +
		tabBar + "\n\n" +
		st.normal.Render(content)

	return st.mainBox.
		Width(width).
		Height(height).
		Render(body)
}
