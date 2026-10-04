package sidepanel

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"mvp-sshowroom/ui2/styles"
)

// View renders the topic list inside the panel border.
//
// Sub-items are indented via Topic.Indent, and the highlighted row is painted
// with that topic's own color as a background. That same color is reused by the
// main panel, so the selected topic's color stays consistent across the screen.
func (m Model) View(width, height int) string {
	// Usable text area once the panel padding is removed.
	innerW := max(width-2*styles.PanelPaddingX, 1)
	innerH := max(height-2*styles.PanelPaddingY, 1)

	lines := make([]string, 0, len(m.topics))

	for i, topic := range m.topics {
		indent := strings.Repeat("  ", topic.Indent)
		label := indent + topic.Title

		style := m.styles.Item
		if i == m.cursor {
			// Highlight the selected row with the topic color as the *background*
			// (not the text color). OnAccent then picks black or white for the
			// label so it stays readable on any accent. Width makes the highlight
			// span the whole text area rather than just the few letters of the
			// title.
			accent := styles.Accent(topic.Color)
			label = indent + "> " + topic.Title
			style = m.styles.Selected.
				Width(innerW).
				Background(accent).
				Foreground(styles.OnAccent(accent))
		}

		lines = append(lines, style.Render(label))
	}

	// Clip to the text area before the border/padding are added, so a long list
	// never pushes the panel past its allotted height.
	body := lipgloss.NewStyle().
		MaxWidth(innerW).
		MaxHeight(innerH).
		Render(lipgloss.JoinVertical(lipgloss.Left, lines...))

	return m.styles.Panel.
		Width(width).
		Height(height).
		Render(body)
}
