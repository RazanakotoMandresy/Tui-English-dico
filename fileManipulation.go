package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func searchWord(file *os.File, upperWord string) {
	byteValues, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error during reading file", err)
		return
	}
	var data map[string]interface{}
	err = json.Unmarshal(byteValues, &data)
	if err != nil {
		fmt.Println("error unmarshal data ", err)
		return
	}
	element, ok := data[upperWord]
	if !ok {
		fmt.Println("the word you want to describe not found")
		return
	}
	fmt.Printf("the description of %v is %v \n", upperWord, element)
}
