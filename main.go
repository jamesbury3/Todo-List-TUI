package main

import (
	"fmt"
	"os"
	"todo-list/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		model.InitialModel(),
		tea.WithAltScreen(), // Use alternate screen buffer to prevent scrolling
	)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if m, ok := finalModel.(model.Model); ok && m.SaveError() != "" {
		fmt.Fprintf(os.Stderr, "Error: %s\n", m.SaveError())
		os.Exit(1)
	}
}
