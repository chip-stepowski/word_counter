package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	text := `The Green elephANt knew the truth about The Red Otter,but did nOt tell Anyone Else.`

	// lowercase
	lower := strings.ToLower(text)

	// remove punctuation
	splitFunc := func(r rune) bool {
		return unicode.IsSpace(r) || r == ',' || r == '.' || r == ';'
	}
	// returns an slice
	rpunc := strings.FieldsFunc(lower, splitFunc)
	fmt.Println(rpunc)

	// count the total number of words
	totalWords := len(rpunc)
	fmt.Println("Total number of words:", totalWords)

	// create an empty map then loop over it to count occurrences of each word
	wordMap := make(map[string]int)

	for _, word := range rpunc {
		wordMap[word]++
	}

	fmt.Println(wordMap)
}
