package ui

// Activity represents a single piece of content shown as a tab inside the
// main panel for a given Topic (e.g. "Overview", "Tech Stack", "Demo").
type Activity struct {
	Title   string
	Content string
}

// Topic represents an entry in the left sidebar. Each Topic owns its own set
// of Activities, which become the tabs rendered in the main panel whenever
// that Topic is selected.
type Topic struct {
	Name       string
	Activities []Activity
}

// mockTopics returns hard-coded sample data so the TUI has something
// meaningful to render without needing a backend. In a real application this
// would likely be loaded from a database, API, or config file and passed
// into NewModel instead.
func mockTopics() []Topic {
	return []Topic{
		{
			Name: "🚀 Projects",
			Activities: []Activity{
				{
					Title:   "Overview",
					Content: "SSHowroom is a terminal portfolio served directly over SSH.\nNo browser, no JS — just a Bubble Tea program streamed to your client.",
				},
				{
					Title:   "Tech Stack",
					Content: "- Go\n- Bubble Tea (Elm architecture)\n- Lipgloss (styling)\n- Wish (SSH middleware)",
				},
				{
					Title:   "Demo",
					Content: "Try navigating:\n  j/k   -> move between topics\n  h/l   -> switch tabs\n  q     -> quit",
				},
			},
		},
		{
			Name: "👤 About Me",
			Activities: []Activity{
				{
					Title:   "Bio",
					Content: "I'm a Go developer who enjoys building fast, minimal terminal tools.",
				},
				{
					Title:   "Skills",
					Content: "- Distributed systems\n- TUI/CLI tooling\n- SSH & networking\n- Kubernetes",
				},
			},
		},
		{
			Name: "📬 Contact",
			Activities: []Activity{
				{
					Title:   "Email",
					Content: "hello@example.com",
				},
				{
					Title:   "Socials",
					Content: "GitHub:   github.com/example\nLinkedIn: linkedin.com/in/example",
				},
			},
		},
	}
}
