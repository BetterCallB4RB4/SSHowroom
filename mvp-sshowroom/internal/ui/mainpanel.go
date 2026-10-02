package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
)

type mainModel struct {
	selectedCard int
	openURL      string
}

func newMainPanel() mainModel {
	return mainModel{}
}

func (m mainModel) ResetSelection() mainModel {
	m.selectedCard = 0
	m.openURL = ""
	return m
}

func (m mainModel) Update(msg tea.Msg, topic Topic, columns int) mainModel {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok || len(topic.Cards) == 0 {
		return m
	}
	if columns < 1 {
		columns = 1
	}
	if m.selectedCard >= len(topic.Cards) {
		m.selectedCard = len(topic.Cards) - 1
	}

	row, col := m.selectedCard/columns, m.selectedCard%columns
	switch keyMsg.String() {
	case "h", "left":
		if col > 0 {
			m.selectedCard--
		}
	case "l", "right":
		if col < columns-1 && m.selectedCard+1 < len(topic.Cards) {
			m.selectedCard++
		}
	case "k", "up":
		if row > 0 {
			m.selectedCard -= columns
		}
	case "j", "down":
		if target := m.selectedCard + columns; target < len(topic.Cards) {
			m.selectedCard = target
		}
	case "enter":
		m.openURL = topic.Cards[m.selectedCard].URL
	default:
		m.openURL = ""
	}
	return m
}

func (m mainModel) View(st styles, width, height int, topic Topic) string {
	if width < 1 || height < 1 {
		return ""
	}
	contentWidth := max(1, width-2)
	columns := max(1, contentWidth/34)
	if len(topic.Cards) > 0 && columns > len(topic.Cards) {
		columns = len(topic.Cards)
	}
	if columns < 1 {
		columns = 1
	}
	cardWidth := max(1, (contentWidth-(columns-1)*2)/columns)
	if m.selectedCard >= len(topic.Cards) {
		m.selectedCard = max(0, len(topic.Cards)-1)
	}

	var cards []string
	for i, card := range topic.Cards {
		cards = append(cards, renderCard(st, card, cardWidth, i == m.selectedCard))
	}
	var rows []string
	for start := 0; start < len(cards); start += columns {
		end := min(start+columns, len(cards))
		rowCards := cards[start:end]
		for i := range rowCards {
			if i < len(rowCards)-1 {
				rowCards[i] = lipgloss.NewStyle().Width(cardWidth).Render(rowCards[i])
			}
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, rowCards...))
	}
	body := strings.Join(rows, "\n\n")
	if len(topic.Cards) == 0 {
		body = st.info.Render("No cards in this category yet.")
	}
	if m.openURL != "" {
		body += "\n\n" + st.info.Render("Selected link: "+m.openURL+" (click the card link to open)")
	}
	content := lipgloss.JoinVertical(lipgloss.Left, st.panelTitle.Render(topic.Name), "", body)
	return st.panel.Width(width).Height(height).Render(content)
}

func renderCard(st styles, card Card, width int, selected bool) string {
	if width < 1 {
		return ""
	}
	cardFrame := st.card
	if selected {
		cardFrame = st.selectedCard
	}
	innerWidth := max(1, width-4)
	underline := st.cardUnderline.Foreground(card.Underline).Render(strings.Repeat("━", innerWidth))
	url := card.URL
	if url != "" {
		url = ansi.SetHyperlink(url) + st.cardLink.Render(url) + ansi.ResetHyperlink()
	}
	content := lipgloss.JoinVertical(lipgloss.Left,
		st.cardTitle.Render(card.Title),
		underline,
		st.cardDescription.Width(innerWidth).Render(card.Description),
		"",
		url,
	)
	return cardFrame.Width(width).Render(content)
}
