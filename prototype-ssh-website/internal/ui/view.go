package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.Width == 0 || m.Height == 0 {
		return "Initializing..."
	}

	r := m.Styles.Renderer
	selectedKey := m.SidebarItems[m.SidebarIdx]
	section := m.Sections[selectedKey]
	sectionColor := lipgloss.Color(section.Color)
	activeTabIndex := m.TabIndices[selectedKey]

	// Palette Colori
	woodBase := lipgloss.Color("#8B4513")
	woodShadow := lipgloss.Color("#5D2E0C")
	leafBase := lipgloss.Color("#228B22")
	leafShadow := lipgloss.Color("#006400")
	leafHighlight := lipgloss.Color("#32CD32")
	potBase := lipgloss.Color("#DAA520")
	potShadow := lipgloss.Color("#9B7E10")
	feetColor := lipgloss.Color("#B8860B")

	// Dynamic styles
	dynamicSidebarSelected := m.Styles.SidebarSelectedStyle.Copy().Background(sectionColor)
	dynamicContentTitle := m.Styles.ContentTitleStyle.Copy().Foreground(sectionColor)
	dynamicShortcutKey := m.Styles.ShortcutKeyStyle.Copy().Foreground(sectionColor)

	bonsaiLines := strings.Split(m.AsciiArt, "\n")
	artHeight, artWidth := len(bonsaiLines), 0
	for _, l := range bonsaiLines { if len(l) > artWidth { artWidth = len(l) } }
	canvasHeight := artHeight
	if GroundLevelY+1 > canvasHeight { canvasHeight = GroundLevelY + 1 }

	sidebarWidth := artWidth + 8
	if sidebarWidth < 45 { sidebarWidth = 45 }
	innerWidth := sidebarWidth - 6

	type cell struct { char rune; color lipgloss.Color }
	canvas := make([][]cell, canvasHeight)
	for i := range canvas {
		canvas[i] = make([]cell, artWidth)
		for j := range canvas[i] { canvas[i][j] = cell{char: ' '} }
		if i < artHeight {
			for j, char := range bonsaiLines[i] {
				var col lipgloss.Color
				if i == 12 { col = feetColor } else if i >= 9 && i <= 11 {
					if j < 5 || j > 25 { col = potShadow } else { col = potBase }
				} else if strings.ContainsRune("M,.\"", char) {
					if (i+j)%3 == 0 { col = leafHighlight } else if (i+j)%3 == 1 { col = leafBase } else { col = leafShadow }
				} else if char != ' ' {
					if strings.ContainsRune("_.\"", char) || j < 15 { col = woodShadow } else { col = woodBase }
				}
				canvas[i][j] = cell{char: char, color: col}
			}
		}
	}
	for _, leaf := range m.Fallen {
		ly, lx := int(leaf.Y), int(leaf.X)
		if ly >= 0 && ly < canvasHeight && lx >= 0 && lx < artWidth { canvas[ly][lx] = cell{char: leaf.Char, color: leafBase} }
	}
	for _, leaf := range m.Leaves {
		ly, lx := int(leaf.Y), int(leaf.X)
		if ly >= 0 && ly < canvasHeight && lx >= 0 && lx < artWidth { canvas[ly][lx] = cell{char: leaf.Char, color: leafHighlight} }
	}

	var animatedBonsaiBuilder strings.Builder
	for _, row := range canvas {
		for _, c := range row {
			if c.char == ' ' { animatedBonsaiBuilder.WriteRune(' ') } else { animatedBonsaiBuilder.WriteString(r.NewStyle().Foreground(c.color).Render(string(c.char))) }
		}
		animatedBonsaiBuilder.WriteByte('\n')
	}
	asciiArt := r.NewStyle().MarginBottom(1).Render(animatedBonsaiBuilder.String())

	var contactLines []string
	contactTitleStyle := r.NewStyle().Foreground(m.Styles.TsWhite).Bold(true).MarginBottom(0)
	contactLines = append(contactLines, contactTitleStyle.Render("CONTACTS"))
	for _, c := range m.Contacts {
		labelStyle := m.Styles.ShortcutKeyStyle.Copy().Foreground(lipgloss.Color(c.Color))
		valueStyle := r.NewStyle().Foreground(m.Styles.TsWhite)
		contactLines = append(contactLines, fmt.Sprintf("%-5s %s", labelStyle.Render(c.Platform), valueStyle.Render(c.Value)))
	}
	contacts := lipgloss.JoinVertical(lipgloss.Left, contactLines...)

	var menuLines []string
	menuTitleStyle := contactTitleStyle.Copy().MarginTop(1)
	menuLines = append(menuLines, menuTitleStyle.Render("SECTIONS"))
	for i, item := range m.SidebarItems {
		style := m.Styles.SidebarItemStyle
		if m.Focus == FocusSidebar && i == m.SidebarIdx { style = dynamicSidebarSelected }
		menuLines = append(menuLines, style.Width(innerWidth).Render(item))
	}
	menu := lipgloss.JoinVertical(lipgloss.Left, menuLines...)

	var skillLines []string
	skillHeaderStyle := r.NewStyle().Foreground(m.Styles.TsWhite).Bold(true).MarginTop(2)
	skillLines = append(skillLines, skillHeaderStyle.Width(innerWidth).Align(lipgloss.Center).Render("- SKILLS -"), "")
	for _, s := range m.Skills {
		skillNameStyle := r.NewStyle().Foreground(sectionColor).Bold(true).Width(innerWidth).Align(lipgloss.Center)
		skillLines = append(skillLines, skillNameStyle.Render(s.Name))
		labelStyle := r.NewStyle().Foreground(m.Styles.TsWhite)
		leftLabel, rightLabel := labelStyle.Render(s.LeftLabel), labelStyle.Render(s.RightLabel)
		spaceLen := innerWidth - lipgloss.Width(leftLabel) - lipgloss.Width(rightLabel)
		if spaceLen < 1 { spaceLen = 1 }
		skillLines = append(skillLines, leftLabel+strings.Repeat(" ", spaceLen)+rightLabel)
		cursorPos := int(float64(innerWidth-2) * s.Percentage)
		var barBuilder strings.Builder
		for j := 0; j < innerWidth; j++ {
			if j == cursorPos || j == cursorPos+1 { barBuilder.WriteString(r.NewStyle().Foreground(sectionColor).Render("█")) } else { barBuilder.WriteString(r.NewStyle().Foreground(m.Styles.TsGray).Render("░")) }
		}
		skillLines = append(skillLines, barBuilder.String(), "")
	}
	skills := lipgloss.JoinVertical(lipgloss.Left, skillLines...)

	sidebar := m.Styles.SidebarStyle.Width(sidebarWidth).Height(m.Height - 12).Render(lipgloss.JoinVertical(lipgloss.Left, asciiArt, contacts, menu, skills))

	mainWidth := sidebarWidth * 2
	if sidebarWidth+mainWidth+4 > m.Width { mainWidth = m.Width - sidebarWidth - 6 }

	var tabLines []string
	for i, t := range section.Tabs {
		var style lipgloss.Style
		if i == activeTabIndex {
			style = r.NewStyle().Background(sectionColor).Foreground(m.Styles.TsBlack).Bold(true).Padding(0, 1).MarginRight(1)
		} else {
			style = r.NewStyle().Foreground(m.Styles.TsWhite).Padding(0, 1).MarginRight(1)
		}
		tabLines = append(tabLines, style.Render(t.Title))
	}
	tabsHeader := lipgloss.JoinHorizontal(lipgloss.Top, tabLines...)
	
	activeTab := section.Tabs[activeTabIndex]
	content := lipgloss.JoinVertical(lipgloss.Left, dynamicContentTitle.Render(section.Title), "\n", tabsHeader, "\n", m.Styles.ContentBodyStyle.Width(mainWidth-4).Render(activeTab.Content))
	mainPanel := m.Styles.MainBoxStyle.Width(mainWidth).Height(m.Height - 12).Render(content)

	var headerParts []string
	for i, item := range m.HeaderItems {
		style := m.Styles.HeaderItemStyle
		if m.Focus == FocusHeader && i == m.HeaderIdx { style = r.NewStyle().Foreground(sectionColor).Bold(true).Underline(true).MarginRight(2) }
		headerParts = append(headerParts, style.Render(item))
	}
	header := m.Styles.HeaderStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, headerParts...))

	shortcuts := []string{
		fmt.Sprintf("%s %s", dynamicShortcutKey.Render("j/k"), r.NewStyle().Foreground(m.Styles.TsWhite).Render("sections")),
		fmt.Sprintf("%s %s", dynamicShortcutKey.Render("h/l"), r.NewStyle().Foreground(m.Styles.TsWhite).Render("tabs")),
		fmt.Sprintf("%s %s", dynamicShortcutKey.Render("d"), r.NewStyle().Foreground(m.Styles.TsWhite).Render("download CV")),
		fmt.Sprintf("%s %s", dynamicShortcutKey.Render("q"), r.NewStyle().Foreground(m.Styles.TsWhite).Render("quit")),
	}
	footer := m.Styles.FooterStyle.Width(sidebarWidth + mainWidth + 2).Render(lipgloss.JoinHorizontal(lipgloss.Top, strings.Join(shortcuts, "  |  ")))

	uiContent := lipgloss.JoinVertical(lipgloss.Left, header, lipgloss.JoinHorizontal(lipgloss.Top, sidebar, mainPanel), footer)
	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, uiContent)
}
