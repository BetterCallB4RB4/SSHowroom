package ui

import (
	"strings"
	"time"

	"github.com/charmbracelet/harmonica"
	"github.com/charmbracelet/lipgloss"
)

const (
	faceMinWidth = 110
	faceTickRate = 10
	eyeWidth     = 13
	eyeHeight    = 9
)

type faceModel struct {
	spring        harmonica.Spring
	smile         float64
	smileVelocity float64
	smileTarget   float64
	gazeSpring    harmonica.Spring
	gazeX, gazeY  float64
	gazeXVelocity float64
	gazeYVelocity float64
	gazeTargetX   float64
	gazeTargetY   float64
	frame         int
	happyUntil    time.Time
	blinkUntil    time.Time
	nextBlink     time.Time
}

func newFace() faceModel {
	now := time.Now()
	return faceModel{
		spring:     harmonica.NewSpring(harmonica.FPS(faceTickRate), 5, 0.72),
		gazeSpring: harmonica.NewSpring(harmonica.FPS(faceTickRate), 4, 0.8),
		nextBlink:  now.Add(3 * time.Second),
	}
}

func (m faceModel) MakeHappy(now time.Time) faceModel {
	m.happyUntil = now.Add(1400 * time.Millisecond)
	m.smileTarget = 1
	return m
}

func (m faceModel) LookAt(key string, now time.Time) faceModel {
	switch key {
	case "h", "left":
		m.gazeTargetX = -2.2
		m.gazeTargetY = 0
	case "l", "right":
		m.gazeTargetX = 2.2
		m.gazeTargetY = 0
	case "k", "up":
		m.gazeTargetX = 0
		m.gazeTargetY = -1.3
	case "j", "down":
		m.gazeTargetX = 0
		m.gazeTargetY = 1.3
	default:
		m.gazeTargetX = 0
		m.gazeTargetY = 0
	}
	return m.MakeHappy(now)
}

func (m faceModel) Update(now time.Time) faceModel {
	m.frame++
	if now.Before(m.happyUntil) {
		m.smileTarget = 1
	} else {
		m.smileTarget = 0
	}
	m.smile, m.smileVelocity = m.spring.Update(m.smile, m.smileVelocity, m.smileTarget)
	m.gazeX, m.gazeXVelocity = m.gazeSpring.Update(m.gazeX, m.gazeXVelocity, m.gazeTargetX)
	m.gazeY, m.gazeYVelocity = m.gazeSpring.Update(m.gazeY, m.gazeYVelocity, m.gazeTargetY)
	if m.frame%35 == 0 {
		m.gazeTargetX = 0.7 * float64((m.frame/35)%3-1)
		m.gazeTargetY = 0.3 * float64((m.frame/70)%3-1)
	}

	if now.After(m.nextBlink) {
		m.blinkUntil = now.Add(180 * time.Millisecond)
		m.nextBlink = now.Add(time.Duration(3+((m.frame/faceTickRate)%4)) * time.Second)
	}
	return m
}

func (m faceModel) View(st styles, width, height int, now time.Time) string {
	if width < 1 || height < 1 {
		return ""
	}

	happy := m.smile > 0.55
	blink := now.Before(m.blinkUntil)
	leftEye := renderEye(st, m.gazeX, m.gazeY, happy, blink)
	rightEye := renderEye(st, m.gazeX, m.gazeY, happy, blink)
	eyes := lipgloss.JoinHorizontal(lipgloss.Top, leftEye, "   ", rightEye)
	mouth := st.faceMouth.Foreground(st.faceAccent).Render("       .-''''-.       ")
	if happy {
		mouth = st.faceMouth.Foreground(st.faceHappy).Render("       \\____/       ")
	}
	face := lipgloss.JoinVertical(lipgloss.Center, eyes, mouth)
	expression := "idle / looking around"
	if happy {
		expression = "happy to see you!"
	}
	body := lipgloss.Place(width, max(1, height-3), lipgloss.Center, lipgloss.Center, face)
	content := lipgloss.JoinVertical(lipgloss.Center,
		st.panelTitle.Render("SSHOWROOM EYES"),
		body,
		st.info.Render(expression),
	)
	return st.panel.Width(width).Height(height).Render(content)
}

func renderEye(st styles, gazeX, gazeY float64, happy, blink bool) string {
	h := eyeHeight
	if happy {
		h = 7
	}
	if blink {
		h = 3
	}
	centerX, centerY := (eyeWidth-1)/2, (h-1)/2
	irisX := centerX + int(gazeX+0.5)
	irisY := centerY + int(gazeY+0.5)
	var out strings.Builder
	for y := 0; y < h; y++ {
		for x := 0; x < eyeWidth; x++ {
			dx := float64(x-centerX) / float64(centerX)
			dy := float64(y-centerY) / float64(centerY)
			cell := " "
			switch {
			case dx*dx+dy*dy > 1:
				cell = " "
			case blink:
				cell = st.faceLid.Render("▄")
			case (x-irisX)*(x-irisX)+(y-irisY)*(y-irisY) <= 1:
				cell = st.facePupil.Render(" ")
			case (x-irisX)*(x-irisX)+(y-irisY)*(y-irisY) <= 6:
				cell = st.faceIris.Render(" ")
			default:
				cell = st.faceWhite.Render(" ")
			}
			out.WriteString(cell)
		}
		if y < h-1 {
			out.WriteByte('\n')
		}
	}
	return out.String()
}
