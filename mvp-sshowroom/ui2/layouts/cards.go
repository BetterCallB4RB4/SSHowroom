package layouts

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"mvp-sshowroom/ui2/data"
	"mvp-sshowroom/ui2/styles"
)

// Cards renders a responsive grid of transparent, shadowed cards.
type Cards struct{}

const (
	// minCardWidth is the narrowest a card may get before we use fewer columns.
	minCardWidth = 24
	// cardGap is the blank space inserted between columns.
	cardGap = 2
)

// Render arranges the topic's cards into as many columns as the width allows.
//
// The column count is derived from the width rather than hardcoded, so the same
// layout works on a narrow PTY and a wide one. Cards flow left-to-right and
// wrap onto new rows, with a blank line between rows so the shadows breathe.
func (Cards) Render(width, height int, topic data.Topic, content data.Content, st styles.Styles) string {
	cards := content.Cards[topic.ID]
	if len(cards) == 0 {
		return st.Muted.Render("No cards for this topic yet.")
	}

	accent := styles.Accent(topic.Color)

	// Integer division picks how many (minCardWidth + gap) blocks fit.
	cols := max(width/(minCardWidth+cardGap), 1)
	cardWidth := (width - (cols-1)*cardGap) / cols

	// The box is one cell narrower than the column so its right-edge shadow
	// still fits inside the column width.
	boxWidth := max(cardWidth-1, 1)
	inner := max(boxWidth-2*styles.CardPaddingX, 1)
	gap := strings.Repeat(" ", cardGap)

	var rows []string
	for i := 0; i < len(cards); i += cols {
		end := min(i+cols, len(cards))

		bodies := make([]string, 0, cols)
		for _, card := range cards[i:end] {
			bodies = append(bodies, renderCardBody(card, inner, accent, st))
		}

		// Pad every card in the row to the tallest body so their bottom drop
		// shadows line up instead of ending at different rows.
		tallest := 0
		for _, body := range bodies {
			tallest = max(tallest, lineCount(body))
		}

		rendered := make([]string, 0, cols*2)
		for j, body := range bodies {
			if j > 0 {
				rendered = append(rendered, gap)
			}
			rendered = append(rendered, renderCard(boxWidth, padLines(body, tallest), st))
		}

		if len(rows) > 0 {
			rows = append(rows, "") // vertical gap between card rows
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, rendered...))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

// renderCardBody builds the text of one card: a title with a topic-color
// underline, a description and an external link. It has no background so the
// card body stays transparent.
func renderCardBody(card data.Card, inner int, accent lipgloss.Color, st styles.Styles) string {
	parts := []string{
		// Underline inherits the foreground color, so this is the colored
		// underline required by the spec.
		st.Title.Foreground(accent).Underline(true).Width(inner).Render(card.Title),
		"",
		st.Item.Width(inner).Render(card.Description),
	}
	if card.Link != "" {
		parts = append(parts, "", st.Link.Foreground(accent).Render(card.Link))
	}
	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

// renderCard wraps a body in the card style, which adds padding and the
// right/bottom drop shadow.
func renderCard(width int, body string, st styles.Styles) string {
	return st.Card.Width(width).Render(body)
}

// lineCount returns the number of lines in s.
func lineCount(s string) int {
	return strings.Count(s, "\n") + 1
}

// padLines appends blank lines so s has at least n lines.
func padLines(s string, n int) string {
	if have := lineCount(s); have < n {
		return s + strings.Repeat("\n", n-have)
	}
	return s
}
