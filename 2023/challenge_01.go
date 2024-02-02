package main

import (
	"fmt"
	"os"
	"strings"
)

func Challenge_01() {
	data, err := os.ReadFile("message_01.txt")
	if err != nil {
		panic(err)
	}

	words := strings.Fields(strings.ToLower(string(data)))

	wordMap := make(map[string]int)
	var uniqueWords []string

	for _, word := range words {
		if _, exists := wordMap[word]; !exists {
			wordMap[word] = 1
			uniqueWords = append(uniqueWords, word)
		} else {
			wordMap[word]++
		}
	}

	var result strings.Builder
	for _, word := range uniqueWords {
		result.WriteString(fmt.Sprintf("%s%d", word, wordMap[word]))
	}

	fmt.Printf("%s", result.String())
}
