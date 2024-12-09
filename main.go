package main

import (
	"log"

	"github.com/RazanakotoMandresy/Tui-English-dico/pkg/getWord"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO adding the way to add word
func main() {
	p := tea.NewProgram(getWord.InialGet())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
