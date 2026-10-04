// Package data loads and models the portfolio content shown by the TUI.
//
// The types here map one-to-one onto the YAML files in ui2/assets. Keeping the
// models in a dedicated package means both the loader and the components can
// agree on the shape of the data without importing each other.
package data

// Topic is one row in the side panel.
type Topic struct {
	// ID is the unique key used to look up the topic's cards/CV entries.
	ID string `yaml:"id"`
	// Title is shown in the side list and as the main-panel heading.
	Title string `yaml:"title"`
	// Color is the topic's highlight color, reused across the whole UI.
	Color string `yaml:"color"`
	// Indent is the nesting level: 0 for a top-level item, 1+ for sub-items.
	Indent int `yaml:"indent"`
	// Layout names the main-panel strategy ("cards", "cv", ...).
	Layout string `yaml:"layout"`
}

// Card is one item in a "cards" layout.
type Card struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Link        string `yaml:"link"`
}

// CVEntry is one item in a "cv" layout.
type CVEntry struct {
	Period string `yaml:"period"`
	Role   string `yaml:"role"`
	Org    string `yaml:"org"`
	Detail string `yaml:"detail"`
}

// Content is the fully loaded dataset for one session. Cards and CV are keyed
// by topic id so a topic only ever carries the data that belongs to it.
type Content struct {
	Topics []Topic
	Cards  map[string][]Card
	CV     map[string][]CVEntry
}
