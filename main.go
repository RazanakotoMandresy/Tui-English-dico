package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

// TODO refactoring the code
func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
