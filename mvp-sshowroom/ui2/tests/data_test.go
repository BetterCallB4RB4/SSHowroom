package tests

import (
	"testing"

	"mvp-sshowroom/ui2/data"
)

// TestLoad verifies the embed directive and the YAML shape agree: the files are
// found, decode without error, and every topic resolves to a non-empty body.
// This catches the two easy mistakes -- a renamed file or a misplaced YAML key.
func TestLoad(t *testing.T) {
	content, err := data.Load()
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if len(content.Topics) == 0 {
		t.Fatal("no topics loaded")
	}

	for _, topic := range content.Topics {
		switch topic.Layout {
		case "cards":
			if len(content.Cards[topic.ID]) == 0 {
				t.Errorf("topic %q has no cards", topic.ID)
			}
		case "cv":
			if len(content.CV[topic.ID]) == 0 {
				t.Errorf("topic %q has no CV entries", topic.ID)
			}
		default:
			t.Errorf("topic %q has unknown layout %q", topic.ID, topic.Layout)
		}
	}
}
