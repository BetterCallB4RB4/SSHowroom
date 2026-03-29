package ui

import (
	"math/rand"
	"strings"
	"time"

	"github.com/alexcloudborn/ssh-portfolio/internal/data"
	"github.com/alexcloudborn/ssh-portfolio/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type FocusArea int

const (
	FocusHeader FocusArea = iota
	FocusSidebar
)

type Point struct {
	X int
	Y int
}

type Leaf struct {
	X          float64
	Y          float64
	VY         float64
	Amplitude  float64
	Frequency  float64
	DirectionX float64
	StartTime  time.Time
	Char       rune
}

type Model struct {
	Width        int
	Height       int
	Focus        FocusArea
	HeaderIdx    int
	SidebarIdx   int
	TabIndices   map[string]int
	HeaderItems  []string
	SidebarItems []string
	Sections     map[string]data.Section
	Skills       []data.Skill
	Contacts     []data.Contact
	AsciiArt     string
	Styles       styles.Styles

	// Animation state
	Leaves          []Leaf
	LastLeafAt      time.Time
	LeafSpawnPoints []Point
	
	// Pile and collision state
	Pile     map[int]int
	Fallen   []Leaf
	Occupied map[string]bool
}

func InitialModel(s styles.Styles) Model {
	d := data.GetInitialData()

	// Randomize Skill values for each new login/session
	randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d.Skills {
		// Set a new random percentage between 0.3 and 0.9 for each login
		d.Skills[i].Percentage = 0.3 + randomSource.Float64()*0.6
	}

	var spawnPoints []Point
	lines := strings.Split(d.AsciiArt, "\n")
	for y, line := range lines {
		for x, char := range line {
			if char == 'M' {
				spawnPoints = append(spawnPoints, Point{X: x, Y: y})
			}
		}
	}

	return Model{
		Focus:           FocusSidebar,
		HeaderIdx:       0,
		SidebarIdx:      0,
		TabIndices:      make(map[string]int),
		HeaderItems:     d.HeaderItems,
		SidebarItems:    d.SidebarItems,
		Sections:        d.Sections,
		Skills:          d.Skills,
		Contacts:        d.Contacts,
		AsciiArt:        d.AsciiArt,
		Styles:          s,
		Leaves:          make([]Leaf, 0),
		LeafSpawnPoints: spawnPoints,
		Pile:            make(map[int]int),
		Fallen:          make([]Leaf, 0),
		Occupied:        make(map[string]bool),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		animateTick(),
		spawnTick(),
	)
}

func animateTick() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return animateMsg(t)
	})
}

func spawnTick() tea.Cmd {
	return tea.Tick(time.Second*5, func(t time.Time) tea.Msg {
		return spawnMsg(t)
	})
}

type animateMsg time.Time
type spawnMsg time.Time
