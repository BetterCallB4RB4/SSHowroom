package mainpanel

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"mvp-sshowroom/ui2/bigfont"
	"mvp-sshowroom/ui2/layouts"
	"mvp-sshowroom/ui2/styles"
)

// View renders the panel: a large ASCII-art topic title, then the topic's
// layout.
//
// width/height are the panel's outer *content* dimensions (they include the
// panel padding). Styles.Panel then adds the border around them, and the text
// area is reduced by the padding so content never touches the border.
func (m Model) View(width, height int) string {
	if !m.hasTopic {
		return m.styles.Panel.
			Width(width).
			Height(height).
			Render(m.styles.Muted.Render("Loading..."))
	}

	// Usable text area once the panel padding is removed.
	innerW := max(width-2*styles.PanelPaddingX, 1)
	innerH := max(height-2*styles.PanelPaddingY, 1)

	accent := styles.Accent(m.topic.Color)

	// Draw the topic title as ASCII art, sized to fit the panel. It is painted
	// in the topic color so the highlight still propagates to the heading.
	art := bigfont.Render(m.topic.Title, innerW)
	title := m.styles.Title.Foreground(accent).Render(art)
	titleHeight := strings.Count(art, "\n") + 1

	// Leave room for the title art and a blank line before the body.
	bodyHeight := max(innerH-titleHeight-1, 1)
	body := layouts.For(m.topic.Layout).
		Render(innerW, bodyHeight, m.topic, m.content, m.styles)

	content := lipgloss.JoinVertical(lipgloss.Left, title, "", body)

	// Clip to the text area before the border/padding are added, so an
	// over-long body cannot grow the panel past its allotted size.
	content = lipgloss.NewStyle().
		MaxWidth(innerW).
		MaxHeight(innerH).
		Render(content)

	return m.styles.Panel.
		Width(width).
		Height(height).
		Render(content)
}
