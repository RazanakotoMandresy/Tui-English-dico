package main

type Word struct {
	// le mots qu'on veut decrire
	Meanings map[string][]interface{} `json:"MEANINGS"`
	Antonyms []string                 `json:"ANTONYMS"`
	Synonyms []string                 `json:"SYNONYMS"`
}
