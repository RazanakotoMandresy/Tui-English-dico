package main

import (
	"fmt"
	"os"
	"strings"
)

// TODO Refactoring , adding a ui to make it a TUI 😆 , get all synonymes and make a choise to chose and showing all similar word  if found an dico fench or malgache add them

func check(word string) (*Word, error) {
	upperWord := strings.ToUpper(word)
	if strings.HasPrefix(upperWord, "A") {
		file, err := os.Open("./data/a.json")
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
	if strings.HasPrefix(upperWord, "B") {
		file, err := os.Open("./data/b.json")
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
	if strings.HasPrefix(upperWord, "C") {
		file, err := os.Open("./data/c.json")
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
	if strings.HasPrefix(upperWord, "D") {
		file, err := os.Open("./data/d.json")
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
	if strings.HasPrefix(upperWord, "E") {
		file, err := os.Open("./data/e.json")
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
	if strings.HasPrefix(upperWord, "F") {
		file, err := os.Open("./data/f.json")
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
	if strings.HasPrefix(upperWord, "G") {
		file, err := os.Open("./data/g.json")
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
	if strings.HasPrefix(upperWord, "H") {
		file, err := os.Open("./data/h.json")
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
	if strings.HasPrefix(upperWord, "I") {
		file, err := os.Open("./data/i.json")
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
	if strings.HasPrefix(upperWord, "J") {
		file, err := os.Open("./data/j.json")
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
	if strings.HasPrefix(upperWord, "K") {
		file, err := os.Open("./data/k.json")
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
	if strings.HasPrefix(upperWord, "L") {
		file, err := os.Open("./data/l.json")
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
	if strings.HasPrefix(upperWord, "M") {
		file, err := os.Open("./data/m.json")
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
	if strings.HasPrefix(upperWord, "N") {
		file, err := os.Open("./data/n.json")
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
	if strings.HasPrefix(upperWord, "O") {
		file, err := os.Open("./data/o.json")
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
	if strings.HasPrefix(upperWord, "P") {
		file, err := os.Open("./data/p.json")
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
	if strings.HasPrefix(upperWord, "Q") {
		file, err := os.Open("./data/q.json")
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
	if strings.HasPrefix(upperWord, "R") {
		file, err := os.Open("./data/r.json")
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
	if strings.HasPrefix(upperWord, "S") {
		file, err := os.Open("./data/s.json")
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
	if strings.HasPrefix(upperWord, "T") {
		file, err := os.Open("./data/t.json")
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
	if strings.HasPrefix(upperWord, "U") {
		file, err := os.Open("./data/a.json")
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
	if strings.HasPrefix(upperWord, "V") {
		file, err := os.Open("./data/v.json")
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
	if strings.HasPrefix(upperWord, "W") {
		file, err := os.Open("./data/w.json")
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
	if strings.HasPrefix(upperWord, "X") {
		file, err := os.Open("./data/x.json")
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
	if strings.HasPrefix(upperWord, "Y") {
		file, err := os.Open("./data/y.json")
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
	if strings.HasPrefix(upperWord, "Z") {
		file, err := os.Open("./data/z.json")
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
	// TODO adding the error messages like don't start with letter on alphabet
	return nil, fmt.Errorf("Word should start from A-Z you have enter %v \n ", word)
}
