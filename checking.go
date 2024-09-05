package main

import (
	"fmt"
	"os"
	"strings"
)

func check(word string) (*Word, error) {
	strSlice := strings.Split(word, "")
	upperWord := strings.ToUpper(word)
	if len(strSlice) == 0 {
		return nil, fmt.Errorf("invalid  invalid you enter nothing %v", word)
	}
	if strings.Contains(word, " ") {
		return nil, fmt.Errorf("word to describe cannot contains any empty space  in %v", word)
	}
	fileToOpen := fmt.Sprintf("./data/%v.json", strSlice[0])
	file, err := os.Open(fileToOpen)
	if err != nil {
		return nil, fmt.Errorf("error during opening file %v", err)
	}
	defer file.Close()
	result, err := searchWord(file, upperWord)
	if err != nil {
		return nil, err
	}
	return result, nil
}
