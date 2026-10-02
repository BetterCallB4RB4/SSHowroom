package ui

import (
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const terminalMinWidth = 120

type terminalModel struct {
	tick int
}

func (m terminalModel) Update() terminalModel {
	m.tick++
	return m
}

func (m terminalModel) View(st styles, width, height int) string {
	if width < 1 || height < 1 {
		return ""
	}

	artWidth := min(width-2, 40)
	artHeight := min(height-5, 24)
	if artWidth < 10 || artHeight < 6 {
		return st.panel.Width(width).Height(height).Render(st.panelTitle.Render("GHOSTTIME"))
	}
	frame := renderSpinningGhost(st, m.tick, artWidth, artHeight)
	content := lipgloss.Place(width, max(1, height-3), lipgloss.Center, lipgloss.Center, frame)
	body := lipgloss.JoinVertical(lipgloss.Center,
		st.panelTitle.Render("GHOSTTIME"),
		content,
		st.info.Render("spinning ASCII spirit"),
	)
	return st.panel.Width(width).Height(height).Render(body)
}

func renderSpinningGhost(st styles, tick, width, height int) string {
	const (
		radiusX = 0.78
		radiusY = 0.92
		radiusZ = 0.72
	)
	angle := float64(tick) * 0.045
	cosA, sinA := math.Cos(angle), math.Sin(angle)
	projectedRadius := math.Sqrt(radiusX*radiusX*cosA*cosA + radiusZ*radiusZ*sinA*sinA)
	colors := st.ghostColors
	var out strings.Builder
	for row := 0; row < height; row++ {
		y := (float64(row)/float64(height-1)*2 - 1) * 1.08
		for col := 0; col < width; col++ {
			x := (float64(col)/float64(width-1)*2 - 1) * projectedRadius
			cell, color, filled := ghostCell(x, y, angle, radiusX, radiusY, radiusZ, tick)
			if !filled {
				out.WriteByte(' ')
				continue
			}
			if color >= 0 {
				color %= len(colors)
				if color < 0 {
					color += len(colors)
				}
				out.WriteString(st.ghostOutline.Foreground(colors[color]).Render(string(cell)))
			} else {
				out.WriteString(st.ghostBody.Render(string(cell)))
			}
		}
		if row < height-1 {
			out.WriteByte('\n')
		}
	}
	return out.String()
}

// ghostCell ray-samples a rotating ellipsoid with a scalloped ghost hem.
func ghostCell(screenX, y, angle, rx, ry, rz float64, tick int) (rune, int, bool) {
	bodyY := y + 0.05
	cosA, sinA := math.Cos(angle), math.Sin(angle)
	projectedRadius := math.Sqrt(rx*rx*cosA*cosA + rz*rz*sinA*sinA)
	inside := screenX*screenX/(projectedRadius*projectedRadius)+bodyY*bodyY/(ry*ry) < 1
	if bodyY > 0.48 {
		wave := 0.08 * math.Cos(screenX/projectedRadius*math.Pi*3)
		inside = screenX*screenX/(projectedRadius*projectedRadius)+math.Pow((bodyY-wave)/0.82, 2) < 1
	}
	if !inside {
		return ' ', 0, false
	}
	depthRadius := (rx * rz / projectedRadius) * math.Sqrt(max(0, 1-(screenX/projectedRadius)*(screenX/projectedRadius)-(bodyY/ry)*(bodyY/ry)))
	localX := screenX * rx * rx * cosA / (projectedRadius * projectedRadius)
	localZ := screenX * rz * rz * sinA / (projectedRadius * projectedRadius)
	depth := -localX*sinA + localZ*cosA
	angleAround := math.Atan2(bodyY, screenX)
	rim := depth > depthRadius-0.11 || depth < -depthRadius+0.035 || math.Abs(screenX) > projectedRadius*0.9
	if rim {
		return rimGlyph(angleAround, tick), (tick/2 + int(angleAround*8)) % 6, true
	}

	if depth > 0 && math.Abs(angle) < 1.25 {
		eyeY := bodyY + 0.10
		eyeXLeft := -0.27*cosA + rz*sinA
		eyeXRight := 0.27*cosA + rz*sinA
		if math.Abs(eyeY) < 0.045 && (math.Abs(screenX-eyeXLeft) < 0.075 || math.Abs(screenX-eyeXRight) < 0.075) {
			return '●', -1, true
		}
		if bodyY > 0.20 && bodyY < 0.34 && math.Abs(screenX-rz*sinA) < 0.18 {
			return 'ᴗ', -1, true
		}
	}

	shade := (1 - depth/max(depthRadius, 0.001)) * 0.5
	glyphs := "@$#S%?*+;:,. "
	index := int(shade * float64(len(glyphs)-1))
	if index < 0 {
		index = 0
	}
	if index >= len(glyphs) {
		index = len(glyphs) - 1
	}
	return rune(glyphs[index]), -1, true
}

func rimGlyph(angle float64, tick int) rune {
	glyphs := []rune(".-=+*x")
	return glyphs[(int(math.Abs(angle)*3)+tick/3)%len(glyphs)]
}
