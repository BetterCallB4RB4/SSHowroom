package companionpanel

// View draws the empty bordered placeholder.
//
// width/height are the panel's *content* dimensions: Styles.Panel adds the
// border around them, so the final block is width+2 by height+2 cells.
func (m Model) View(width, height int) string {
	return m.styles.Panel.
		Width(width).
		Height(height).
		Render("")
}
