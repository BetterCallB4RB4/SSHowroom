// Package tests holds black-box tests for the ui2 packages. Keeping them in one
// external package means each production package stays free of test files, and
// the tests can only rely on the exported API.
package tests

import (
	"testing"

	"mvp-sshowroom/ui2/app"
)

// TestComputeSumsBackToUsableArea guards the core grid invariant: the panel
// sizes must add back up to the usable area (terminal minus margins and
// gutters). If integer rounding ever leaks, the layout silently grows a
// one-cell gap or overflow -- exactly the bug this math exists to prevent.
func TestComputeSumsBackToUsableArea(t *testing.T) {
	cases := []struct{ w, h int }{
		{120, 40}, // divisible by 3 once chrome is removed
		{121, 41}, // remainder on both axes
		{100, 30}, // classic terminal
		{200, 60}, // wide
		{80, 40},
	}

	for _, tc := range cases {
		s := app.Compute(tc.w, tc.h)

		wantW := tc.w - 2*app.MarginX - app.GutterX
		if got := s.MainOuterW + s.SideOuterW; got != wantW {
			t.Errorf("width %d: main(%d)+side(%d)=%d, want %d",
				tc.w, s.MainOuterW, s.SideOuterW, got, wantW)
		}

		wantSideH := tc.h - 2*app.MarginY - app.GutterY
		if got := s.CompanionOuterH + s.SideOuterH; got != wantSideH {
			t.Errorf("height %d: companion(%d)+side(%d)=%d, want %d",
				tc.h, s.CompanionOuterH, s.SideOuterH, got, wantSideH)
		}

		if want := tc.h - 2*app.MarginY; s.MainOuterH != want {
			t.Errorf("height %d: main outer=%d, want %d", tc.h, s.MainOuterH, want)
		}

		if s.SideOuterW != s.CompanionOuterW {
			t.Errorf("width %d: side(%d) != companion(%d)", tc.w, s.SideOuterW, s.CompanionOuterW)
		}
	}
}
