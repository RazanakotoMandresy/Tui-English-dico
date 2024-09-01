package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type WordStruct struct {
	Meaning map[string]string
}

func searchWord(file *os.File, upperWord string) {
	// valeur bytes du file open
	byteValues, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error during reading file", err)
		return
	}
	// map[string]interface{} type object json dans go
	var data map[string]interface{}
	err = json.Unmarshal(byteValues, &data)
	if err != nil {
		fmt.Println("error unmarshal data ", err)
		return
	}
	// check dans le data le mots correspendant a upperword
	element, ok := data[upperWord]
	if !ok {
		fmt.Println("the word you want to describe not found")
		return
	}
	fmt.Printf("the description of %v is %v \n", upperWord, element)
	fmt.Println("enter another word you want to describe")
	fmt.Println("..................................................")
	fmt.Printf("%t \n", element)
	fmt.Println("..................................................")
}
