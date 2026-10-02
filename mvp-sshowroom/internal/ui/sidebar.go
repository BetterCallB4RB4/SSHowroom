package ui

import tea "github.com/charmbracelet/bubbletea"

// sidebarModel is the state for the left-hand topics panel. It is a small,
// self-contained Elm-style sub-model: it only knows about its own cursor
// position and how to render itself, nothing about the main panel.
type sidebarModel struct {
	topics []Topic
	cursor int // index of the currently highlighted topic
}

// newSidebar builds the initial sidebar state from the given topics.
func newSidebar(topics []Topic) sidebarModel {
	return sidebarModel{topics: topics}
}

// Selected returns the Topic currently highlighted in the sidebar. The main
// panel uses this to decide which Activities (tabs) to display.
func (s sidebarModel) Selected() Topic {
	return s.topics[s.cursor]
}

func (s sidebarModel) Move(delta int) sidebarModel {
	if len(s.topics) == 0 {
		return s
	}
	s.cursor = (s.cursor + delta + len(s.topics)) % len(s.topics)
	return s
}

// Update handles sidebar-specific navigation. Movement wraps around (moving
// past the last topic jumps back to the first, and vice versa) which is a
// common UX pattern for short vertical lists.
//
// It returns the bool changed so the caller (root Model) knows whether the
// selected topic changed, which lets it reset the main panel's active tab.
func (s sidebarModel) Update(msg tea.Msg) (sidebarModel, bool) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return s, false
	}

	switch keyMsg.String() {
	case "k", "up":
		if s.cursor > 0 {
			s.cursor--
		} else {
			s.cursor = len(s.topics) - 1
		}
		return s, true

	case "j", "down":
		if s.cursor < len(s.topics)-1 {
			s.cursor++
		} else {
			s.cursor = 0
		}
		return s, true
	}

	return s, false
}

// View renders the sidebar into a box of the given width/height. The active
// topic is highlighted with the "selected" style; all others use "normal".
func (s sidebarModel) View(st styles, width, height int) string {
	var body string
	for i, topic := range s.topics {
		line := st.normal.Render(topic.Name)
		if i == s.cursor {
			line = st.selected.Width(max(0, width-4)).Render(topic.Name)
		}
		body += line + "\n"
	}

	return st.panel.
		Width(width).
		Height(height).
		Render(st.panelTitle.Render("Topics") + "\n\n" + body)
}
