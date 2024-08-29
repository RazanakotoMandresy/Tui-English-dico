package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// func mijery hoe start avy am'i inona ila izy misy params word na upperWord ihany izy lo eto mijery de a Z
// Lot of hard coding fix this later 🥲😭

func check(word string) {
	upperWord := strings.ToUpper(word)
	if strings.HasPrefix(upperWord, "A") {
		// Mitady anle mot ao anaty data Json specifique
		file, err := os.Open("./data/a.json")
		if err != nil {
			fmt.Println("error during opening file", err)
			return
		}
		defer file.Close()
		byteValues, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("error during reading file", err)
			return
		}
		// map data type map[string]interface{}
		var data map[string]interface{}
		err = json.Unmarshal(byteValues, &data)
		if err != nil {
			fmt.Println("error unmarshal data ", err)
			return
		}
		element, ok := data[upperWord]
		fmt.Printf("element %v \n , ok %v \n" , element, ok)
	}

	if strings.HasPrefix(upperWord, "B") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "C") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "D") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "E") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "F") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "I") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "G") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "H") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "I") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "J") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "K") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "L") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "M") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "N") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "O") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "P") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "Q") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "R") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "S") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "T") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "U") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "V") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "W") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "X") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
	if strings.HasPrefix(upperWord, "Y") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)

	}
	if strings.HasPrefix(upperWord, "Z") {
		// Mitady anle mot ao anaty data Json specifique
		fmt.Println(upperWord)
	}
}
