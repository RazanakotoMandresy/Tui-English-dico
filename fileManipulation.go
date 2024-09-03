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

func searchWord(file *os.File, upperWord string) *Word {
	// valeur bytes du file open
	byteValues, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error during reading file", err)
		return nil
	}
	// map[string]interface{} type object json dans go
	var data map[string]Word
	err = json.Unmarshal(byteValues, &data)
	if err != nil {
		fmt.Println("error unmarshal data ", err)
		return nil
	}
	// check dans le data le mots correspendant a upperword
	element, ok := data[upperWord]
	if !ok {
		fmt.Println("the word you want to describe not found ")
		return nil
	}
	return &Word{
		Synonyms: element.Synonyms,
		Antonyms: element.Antonyms,
		Meanings: element.Meanings,
	}
}
