// Command ui2 is a standalone development runner for the SSHowroom TUI.
//
// The production SSH server (sshserver/) serves this UI over SSH. This entry
// point exists so the layout can also be developed locally, without SSH, with:
//
//	go run ./ui2
//
// without touching anything outside the ui2/ folder. Wiring ui2 into the SSH
// server later means importing mvp-sshowroom/ui2/app from its handler.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"mvp-sshowroom/ui2/app"
)

func main() {
	program := tea.NewProgram(
		app.New(lipgloss.DefaultRenderer()),
		tea.WithAltScreen(),
	)
	if _, err := program.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "ui2:", err)
		os.Exit(1)
	}
}
