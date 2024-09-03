package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Word struct {
	Antonyms []string    `json:"ANTONYMS"`
	Synonyms []string    `json:"SYNONYMS"`
	Meanings interface{} `json:"MEANINGS"`
}

func searchWord(file *os.File, upperWord string) (*Word, error) {
	// valeur bytes du file open
	byteValues, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error during reading file %v ", err)
	}
	// map[string]interface{} type object json dans go
	var data map[string]Word
	err = json.Unmarshal(byteValues, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal data %v ", err)
	}
	// check dans le data le mots correspendant a upperword
	element, ok := data[upperWord]
	if !ok {
		return nil, fmt.Errorf("the word you want to describe not found")
	}
	return &Word{
		Synonyms: element.Synonyms,
		Antonyms: element.Antonyms,
		Meanings: element.Meanings,
	}, nil
}
