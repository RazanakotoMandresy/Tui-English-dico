package main

import (
	"fmt"
	"os"
	"strings"
)

func check(word string) (*Word, error) {
	strSlice := strings.Split(word, "")
	if len(strSlice) == 0 {
	return nil, fmt.Errorf("invalid input %v", word)
	}
	upperWord := strings.ToUpper(word)
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
