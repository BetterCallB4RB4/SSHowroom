package ui

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

// Layout constants controlling how screen space is divided between the
// panels. Tweak these to change the proportions.
const (
	sidebarWidthRatio = 0.25 // sidebar takes 25% of total width
	footerHeight      = 1    // reserved row for the key-hint footer
	panelGap          = 1    // space between panels
)

const splashTitle = `##### ##### #...# .###. #...# ####. .###. .###. #...#
#.... #.... #...# #...# #...# #...# #...# #...# ##.##
##### ##### ##### #...# #.#.# ####. #...# #...# #.#.#
....# ....# #...# #...# ##.## #.#.. #...# #...# #...#
##### ##### #...# .###. #...# #..## .###. .###. #...#`

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
	if m.splash {
		return m.splashView()
	}

	// Reserve room for the footer and account for the border characters.
	const borderSize = 2
	contentHeight := m.height - footerHeight
	if contentHeight < 0 {
		contentHeight = 0
	}

	showTerminal := m.width >= terminalMinWidth && m.height >= 18
	showFace := m.width >= faceMinWidth && (!showTerminal || m.width >= 190)
	columnGaps := panelGap
	if showFace {
		columnGaps += panelGap
	}
	if showTerminal {
		columnGaps += panelGap
	}
	sidebarWidth := int(float64(m.width) * sidebarWidthRatio)
	mainWidth := m.width - sidebarWidth - columnGaps
	faceWidth := 0
	terminalWidth := 0
	if showFace {
		faceWidth = min(40, max(34, m.width/4))
		mainWidth -= faceWidth
	}
	if showTerminal {
		terminalWidth = min(50, max(40, m.width/4))
		mainWidth -= terminalWidth
	}

	// Subtract border size so the rendered box matches its allotted space.
	sidebarView := m.sidebar.View(m.styles, sidebarWidth-borderSize, contentHeight-borderSize)
	mainView := m.main.View(m.styles, mainWidth-borderSize, contentHeight-borderSize, m.sidebar.Selected())

	columns := []string{sidebarView, lipgloss.NewStyle().Width(panelGap).Render(""), mainView}
	if showFace {
		columns = append(columns,
			lipgloss.NewStyle().Width(panelGap).Render(""),
			m.face.View(m.styles, faceWidth-borderSize, contentHeight-borderSize, time.Now()),
		)
	}
	if showTerminal {
		columns = append(columns,
			lipgloss.NewStyle().Width(panelGap).Render(""),
			m.terminal.View(m.styles, terminalWidth-borderSize, contentHeight-borderSize),
		)
	}
	panels := lipgloss.JoinHorizontal(lipgloss.Top, columns...)
	footer := m.styles.info.Foreground(m.footerColor).Render(" tab: category • hjkl/arrows: cards • enter: link • q: quit ")

	return panels + "\n" + footer
}

func (m Model) splashView() string {
	portrait := renderArt(portraitArt, m.width, max(1, m.height-8))
	title := renderArt(splashTitle, m.width, 5)
	content := lipgloss.JoinVertical(lipgloss.Center,
		portrait,
		title,
		"50% tech guy, 50% techno guy.",
		"Press any key to enter",
	)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
