package ui

import "github.com/charmbracelet/lipgloss"

// Layout constants controlling how screen space is divided between the two
// panels. Tweak these to change the proportions.
const (
	sidebarWidthRatio = 0.3 // sidebar takes 30% of total width
	footerHeight      = 1   // reserved row for the key-hint footer
	panelGap          = 1   // space between sidebar and main panel
)

// View renders the full two-panel layout: sidebar on the left, main content
// panel on the right, with a footer hint bar beneath both.
func (m Model) View() string {
	if m.quitting {
		return "\n  Thanks for visiting SSHowroom! Goodbye!\n\n"
	}

	// Avoid rendering garbage before the first tea.WindowSizeMsg arrives.
	if !m.ready {
		return "\n  Initializing SSHowroom...\n"
	}

	// Reserve room for the footer and account for the 2 columns/rows each
	// bordered box consumes (1 border character on each side).
	const borderSize = 2
	contentHeight := m.height - footerHeight
	if contentHeight < 0 {
		contentHeight = 0
	}

	sidebarWidth := int(float64(m.width) * sidebarWidthRatio)
	mainWidth := m.width - sidebarWidth - panelGap

	// Subtract border size so the *rendered* box matches the space we
	// actually allotted, rather than overflowing it.
	sidebarView := m.sidebar.View(m.styles, sidebarWidth-borderSize, contentHeight-borderSize)
	mainView := m.main.View(m.styles, mainWidth-borderSize, contentHeight-borderSize, m.sidebar.Selected())

	panels := lipgloss.JoinHorizontal(lipgloss.Top, sidebarView, lipgloss.NewStyle().Width(panelGap).Render(""), mainView)

	footer := m.styles.info.Foreground(m.footerColor).Render(" j/k: topics • h/l: tabs • q: quit ")

	return panels + "\n" + footer
}
