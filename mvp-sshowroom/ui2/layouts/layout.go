// Package layouts implements the per-topic main-panel view strategies.
//
// A topic's `layout` field (from topics.yaml) names a strategy. The main panel
// asks the registry for the matching Layout and calls Render, so adding a new
// content style means adding one file here plus one entry in registry -- no
// changes to the panels or the root model are needed.
package layouts

import (
	"mvp-sshowroom/ui2/data"
	"mvp-sshowroom/ui2/styles"
)

// Layout renders the body of the main panel for a single topic.
//
// width and height are the *content* dimensions of the panel (the border is
// already excluded), which lets a strategy arrange itself to the space it has.
type Layout interface {
	Render(width, height int, topic data.Topic, content data.Content, st styles.Styles) string
}

// registry maps a topic's `layout` value to its strategy. A plain map is enough
// because the set is small and fixed at compile time.
var registry = map[string]Layout{
	"cards": Cards{},
	"cv":    CV{},
}

// For returns the strategy registered under name. It falls back to Cards so an
// unknown or misspelled layout still renders something useful instead of
// leaving the main panel blank.
func For(name string) Layout {
	if layout, ok := registry[name]; ok {
		return layout
	}
	return Cards{}
}
