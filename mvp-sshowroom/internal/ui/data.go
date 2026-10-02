package ui

import (
	"math/rand"

	"github.com/charmbracelet/lipgloss"
)

// Card is a portfolio item rendered in the main panel.
type Card struct {
	Title       string
	Description string
	URL         string
	Underline   lipgloss.Color
}

// Topic is a category in the sidebar and owns the cards shown when selected.
type Topic struct {
	Name  string
	Cards []Card
}

// mockTopics provides sample portfolio content without requiring a backend.
func mockTopics() []Topic {
	colors := []lipgloss.Color{
		lipgloss.Color("#FFD75F"),
		lipgloss.Color("#67D9E8"),
		lipgloss.Color("#FF78B7"),
		lipgloss.Color("#78E0A1"),
		lipgloss.Color("#B38CFF"),
		lipgloss.Color("#FF8B65"),
	}
	card := func(title, description, url string) Card {
		underline := colors[rand.Intn(len(colors))]
		return Card{Title: title, Description: description, URL: url, Underline: underline}
	}

	return []Topic{
		{
			Name: "Projects",
			Cards: []Card{
				card("SSHowroom", "A terminal portfolio served directly over SSH, built with Go and Bubble Tea.", "https://google.com"),
				card("TUI Toolkit", "A collection of fast, keyboard-driven terminal interfaces.", "https://google.com"),
				card("Infrastructure", "Deployment and automation experiments for reliable services.", "https://google.com"),
				card("Open Source", "Small tools and contributions focused on practical developer workflows.", "https://google.com"),
			},
		},
		{
			Name: "About Me",
			Cards: []Card{
				card("Bio", "A Go developer who enjoys building fast, minimal terminal tools.", "https://google.com"),
				card("Distributed Systems", "Experience designing and operating resilient services.", "https://google.com"),
				card("TUI and CLI", "Keyboard-first interfaces with a focus on clarity and speed.", "https://google.com"),
				card("Networking", "SSH, service connectivity, and practical network tooling.", "https://google.com"),
			},
		},
		{
			Name: "Contact",
			Cards: []Card{
				card("Email", "Get in touch about projects, collaboration, or just to say hello.", "https://google.com"),
				card("GitHub", "Explore source code, experiments, and open-source work.", "https://google.com"),
				card("LinkedIn", "Connect for professional conversations and opportunities.", "https://google.com"),
			},
		},
	}
}
