package ui

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const GroundLevelY = 13 

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case spawnMsg:
		if len(m.LeafSpawnPoints) > 0 {
			p := m.LeafSpawnPoints[rand.Intn(len(m.LeafSpawnPoints))]
			newLeaf := Leaf{
				X:          float64(p.X),
				Y:          float64(p.Y),
				VY:         0.05 + rand.Float64()*0.1, 
				Amplitude:  0.5 + rand.Float64()*1.5,
				Frequency:  0.5 + rand.Float64()*1.0,
				DirectionX: 1.0,
				StartTime:  time.Time(msg),
				Char:       'W',
			}
			m.Leaves = append(m.Leaves, newLeaf)
		}
		return m, spawnTick()

	case animateMsg:
		var activeLeaves []Leaf
		for i := range m.Leaves {
			leaf := &m.Leaves[i]
			elapsed := time.Since(leaf.StartTime).Seconds()
			nextX := leaf.X + (math.Sin(elapsed*leaf.Frequency)*(leaf.Amplitude/10.0))*leaf.DirectionX
			nextY := leaf.Y + leaf.VY
			ix, iy := int(nextX), int(nextY)
			oldX, oldY := int(leaf.X), int(leaf.Y)

			if ix != oldX && m.isOccupied(ix, oldY) {
				leaf.DirectionX = -leaf.DirectionX
				nextX = leaf.X 
				ix = oldX
			}

			if iy != oldY && (iy >= GroundLevelY || m.isOccupied(oldX, iy)) {
				m.occupy(oldX, oldY)
				m.Fallen = append(m.Fallen, Leaf{X: float64(oldX), Y: float64(oldY), Char: leaf.Char})
				continue 
			}
			leaf.X = nextX
			leaf.Y = nextY
			activeLeaves = append(activeLeaves, *leaf)
		}
		m.Leaves = activeLeaves
		return m, animateTick()

	case tea.KeyMsg:
		selectedKey := m.SidebarItems[m.SidebarIdx]
		section := m.Sections[selectedKey]

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			if m.Focus == FocusHeader {
				m.Focus = FocusSidebar
			} else {
				m.Focus = FocusHeader
			}
		case "j", "down":
			if m.Focus == FocusSidebar {
				m.SidebarIdx = (m.SidebarIdx + 1) % len(m.SidebarItems)
			}
		case "k", "up":
			if m.Focus == FocusSidebar {
				m.SidebarIdx = (m.SidebarIdx - 1 + len(m.SidebarItems)) % len(m.SidebarItems)
			}
		case "l", "right":
			if len(section.Tabs) > 0 {
				m.TabIndices[selectedKey] = (m.TabIndices[selectedKey] + 1) % len(section.Tabs)
			}
		case "h", "left":
			if len(section.Tabs) > 0 {
				m.TabIndices[selectedKey] = (m.TabIndices[selectedKey] - 1 + len(section.Tabs)) % len(section.Tabs)
			}
		}
	}
	return m, nil
}

func (m *Model) isOccupied(x, y int) bool {
	key := fmt.Sprintf("%d,%d", x, y)
	return m.Occupied[key]
}

func (m *Model) occupy(x, y int) {
	key := fmt.Sprintf("%d,%d", x, y)
	m.Occupied[key] = true
}
