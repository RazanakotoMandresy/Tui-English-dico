package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

// "fmt"

// app tui asina ui kely amin'i charm ui semaine prochaine ,
// app dico anglais manome ny sens an'i mots
// manome mot le olona de apres Mitady ao uamin'i data , cas ohatra manomboka amin'i a mijery ao amin'i a.json

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
