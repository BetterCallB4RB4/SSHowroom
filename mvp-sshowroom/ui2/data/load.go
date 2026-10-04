package data

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"mvp-sshowroom/ui2/assets"
)

// File names inside the embedded assets package. Declaring them once avoids
// typos drifting between the embed directive and the decode calls.
const (
	topicsFile = "topics.yaml"
	cardsFile  = "cards.yaml"
	cvFile     = "cv.yaml"
)

// topicsDoc mirrors the top-level shape of topics.yaml (`topics: [...]`).
// The cards and CV files are already plain maps keyed by topic id, so they can
// be decoded straight into the Content fields.
type topicsDoc struct {
	Topics []Topic `yaml:"topics"`
}

// Load reads the embedded YAML files and returns the assembled Content.
//
// It is deliberately synchronous and returns a plain error. The caller wraps it
// in a Bubble Tea command so the decode happens off the update path, keeping
// the Elm runtime responsive even if the content grows.
func Load() (Content, error) {
	var content Content

	topicsRaw, err := assets.Files.ReadFile(topicsFile)
	if err != nil {
		return Content{}, fmt.Errorf("read %s: %w", topicsFile, err)
	}
	var topics topicsDoc
	if err := yaml.Unmarshal(topicsRaw, &topics); err != nil {
		return Content{}, fmt.Errorf("decode %s: %w", topicsFile, err)
	}
	content.Topics = topics.Topics

	cardsRaw, err := assets.Files.ReadFile(cardsFile)
	if err != nil {
		return Content{}, fmt.Errorf("read %s: %w", cardsFile, err)
	}
	if err := yaml.Unmarshal(cardsRaw, &content.Cards); err != nil {
		return Content{}, fmt.Errorf("decode %s: %w", cardsFile, err)
	}

	cvRaw, err := assets.Files.ReadFile(cvFile)
	if err != nil {
		return Content{}, fmt.Errorf("read %s: %w", cvFile, err)
	}
	if err := yaml.Unmarshal(cvRaw, &content.CV); err != nil {
		return Content{}, fmt.Errorf("decode %s: %w", cvFile, err)
	}

	return content, nil
}
