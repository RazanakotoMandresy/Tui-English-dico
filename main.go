package main

import (
	"fmt"
)

// app tui asina ui kely amin'i charm ui semaine prochaine ,
// app dico anglais manome ny sens an'i mots
// manome mot le olona de apres Mitady ao amin'i data , cas ohatra manomboka amin'i a mijery ao amin'i a.json
func main() {
	var word string
	fmt.Println("welcome on TUI-English-dico")
	fmt.Println("enter the words you want to describe")
	_, err := fmt.Scan(&word)
	if err != nil {
		fmt.Println("error when you enter the word to describe")
	}
	check(word)
}
