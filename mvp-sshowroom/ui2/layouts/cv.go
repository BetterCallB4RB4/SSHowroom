package layouts

import (
	"github.com/charmbracelet/lipgloss"

	"mvp-sshowroom/ui2/data"
	"mvp-sshowroom/ui2/styles"
)

// CV renders a plain, vertical list of career/contact entries.
type CV struct{}

// Render prints one block per entry: a muted period, a bold role in the topic
// color, an optional "@ org" suffix, then the detail line.
func (CV) Render(width, height int, topic data.Topic, content data.Content, st styles.Styles) string {
	entries := content.CV[topic.ID]
	if len(entries) == 0 {
		return st.Muted.Render("No entries for this topic yet.")
	}

	accent := styles.Accent(topic.Color)

	blocks := make([]string, 0, len(entries))
	for _, entry := range entries {
		head := st.Title.Foreground(accent).Render(entry.Role)
		if entry.Org != "" {
			head += " " + st.Muted.Render("@ "+entry.Org)
		}

		blocks = append(blocks, lipgloss.JoinVertical(lipgloss.Left,
			st.Muted.Render(entry.Period),
			head,
			st.Item.Width(width).Render(entry.Detail),
		))
	}

	return lipgloss.JoinVertical(lipgloss.Left, interleaveBlanks(blocks)...)
}

// interleaveBlanks puts an empty line between each block so entries breathe.
func interleaveBlanks(blocks []string) []string {
	out := make([]string, 0, len(blocks)*2)
	for i, b := range blocks {
		if i > 0 {
			out = append(out, "")
		}
		out = append(out, b)
	}
	return out
}
