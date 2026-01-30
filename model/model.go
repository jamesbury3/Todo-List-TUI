package model

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

// InitialModel creates and returns the initial model state
func InitialModel() Model {
	// Create backups of all todo files at startup
	err := createBackups()
	if err != nil {
		log.Fatal("Error creating backups: ", err)
	}

	m := Model{
		backlog:     loadTodos(backlogFile),
		ready:       loadTodos(readyFile),
		completed:   loadTodos(completedFile),
		cursor:      0,
		currentView: viewReady,
	}
	m.updateDisplayedCompleted()
	return m
}

// Init initializes the model and returns the initial command
func (m Model) Init() tea.Cmd {
	return tea.ClearScreen
}
