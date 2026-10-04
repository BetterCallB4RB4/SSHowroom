package app

import (
	"github.com/charmbracelet/lipgloss"
)

// View renders the whole screen: the two-column block, centered with margins.
//
// The block is built bottom-up so each join has a known size:
//
//  1. the side column stacks companion (top) over side (bottom), separated by
//     a vertical gutter;
//  2. the block joins the side column (left) and main (right), separated by a
//     horizontal gutter;
//  3. Place centers the block, leaving MarginX/MarginY of empty space around it.
//
// Because the panel sizes are computed from the usable area (terminal minus
// margins and gutters), step 3 centers the block and produces even outer
// margins.
func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		// Bubble Tea has not reported a size yet; render nothing.
		return ""
	}

	if m.err != nil {
		return m.fill("Could not load content:\n" + m.err.Error())
	}

	if TooSmall(m.width, m.height) {
		return m.fill("terminal too small")
	}

	s := Compute(m.width, m.height)

	// A blank block of exactly GutterY rows separates the stacked panels.
	verticalGap := lipgloss.NewStyle().Height(GutterY).Render(" ")
	sideColumn := lipgloss.JoinVertical(
		lipgloss.Top,
		m.companion.View(s.CompanionContentW, s.CompanionContentH),
		verticalGap,
		m.side.View(s.SideContentW, s.SideContentH),
	)

	// A blank block of exactly GutterX columns separates the two columns.
	horizontalGap := lipgloss.NewStyle().Width(GutterX).Render(" ")
	block := lipgloss.JoinHorizontal(
		lipgloss.Top,
		sideColumn,
		horizontalGap,
		m.main.View(s.MainContentW, s.MainContentH),
	)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		block,
	)
}

// fill centers msg in the whole terminal. It is used for the small set of
// non-grid screens (errors, "too small") so they do not need panel styling.
func (m Model) fill(msg string) string {
	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(msg)
}
