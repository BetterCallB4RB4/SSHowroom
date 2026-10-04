package app

// borderSize is the number of columns/rows a lipgloss border adds to a box:
// one cell on the left and one on the right (same for top/bottom).
const borderSize = 2

// Grid spacing, in terminal cells. MarginX/MarginY form the empty border
// around the whole block; GutterX separates the two columns and GutterY
// separates the stacked panels. They are exported so tests and future layout
// tweaks can reason about the geometry.
const (
	MarginX = 4
	MarginY = 1
	GutterX = 4
	GutterY = 1
)

// Minimum unit width and screen height. Below these the three-panel grid stops
// making sense, so View shows a short message instead of mangled output.
const (
	minUnit   = 16
	minHeight = 8
)

// Sizes holds the dimensions of the three panels.
//
// "Outer" is the full block including its border; "Content" is the space left
// for text (outer minus borderSize). Panels are sized with their content
// dimensions and lipgloss adds the border afterwards.
type Sizes struct {
	MainOuterW, MainOuterH               int
	SideOuterW, SideOuterH               int
	CompanionOuterW, CompanionOuterH     int
	MainContentW, MainContentH           int
	SideContentW, SideContentH           int
	CompanionContentW, CompanionContentH int
}

// TooSmall reports whether the terminal is too small to draw the grid,
// including the outer margins and the interior gutters.
func TooSmall(width, height int) bool {
	return width < 3*minUnit+2*MarginX+GutterX ||
		height < minHeight+2*MarginY+GutterY
}

// Compute splits the usable area into the grid described by the layout spec:
//
//	+-----------+  +------------------+
//	| companion |  |                  |  companion: top 1/3 of the side column
//	+-----------+  |                  |
//	|           |  |      main        |  main: 2/3 width, full usable height
//	|   side    |  |                  |
//	|           |  |                  |  side: bottom 2/3 of the side column
//	+-----------+  +------------------+
//
// The whole block is inset by MarginX/MarginY and the columns are separated by
// GutterX, so the usable width is width - 2*MarginX - GutterX. The grid unit is
// one third of that usable width: the side column takes one unit and main takes
// the rest, so the two widths always sum back to the usable width and any
// integer-rounding remainder is absorbed by main instead of leaving a gap.
//
// The same trick on the Y axis splits the usable height into companion
// (height/3) on top and side below; the extra GutterY between them is why main
// is taller than the two side boxes combined.
func Compute(width, height int) Sizes {
	usableW := width - 2*MarginX - GutterX
	usableH := height - 2*MarginY - GutterY

	// The side column takes 30% of the usable width instead of the original
	// third, which gives the main panel a little more room.
	sideOuterW := usableW * 3 / 10
	mainOuterW := usableW - sideOuterW

	companionOuterH := usableH / 3
	sideOuterH := usableH - companionOuterH

	// main spans the usable height plus the internal vertical gutter.
	mainOuterH := usableH + GutterY

	return Sizes{
		MainOuterW: mainOuterW, MainOuterH: mainOuterH,
		SideOuterW: sideOuterW, SideOuterH: sideOuterH,
		CompanionOuterW: sideOuterW, CompanionOuterH: companionOuterH,

		MainContentW: max(mainOuterW-borderSize, 1), MainContentH: max(mainOuterH-borderSize, 1),
		SideContentW: max(sideOuterW-borderSize, 1), SideContentH: max(sideOuterH-borderSize, 1),
		CompanionContentW: max(sideOuterW-borderSize, 1), CompanionContentH: max(companionOuterH-borderSize, 1),
	}
}
