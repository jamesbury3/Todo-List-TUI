package main

import tea "github.com/charmbracelet/bubbletea"

// initialModel creates and returns the initial model state
func initialModel() model {
	m := model{
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
func (m model) Init() tea.Cmd {
	return tea.ClearScreen
}
