package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type charTuple struct {
	char  rune
	index int
}

func Challenge_04() {
	input, err := os.ReadFile("files_quarantine.txt")
	if err != nil {
		panic(err)
	}
	var data = strings.Split(string(input), "\n")

	var count int = 0
	for _, line := range data {
		temp := strings.Split(line, "-")
		str := temp[0]
		hash := temp[1]
		var uniqueChars []charTuple
		var charMap = make(map[rune][2]int)
		for i, char := range str {
			value, exists := charMap[char]
			if !exists {
				value = [2]int{0, 0}
			}

			value[0]++
			value[1] = i

			charMap[char] = value
		}
		for char, c := range charMap {
			if c[0] == 1 {
				tuple := charTuple{char, c[1]}
				uniqueChars = append(uniqueChars, tuple)
			}

		}
		// Sort the slice based on the second member
		sort.Slice(uniqueChars, func(i, j int) bool {
			return uniqueChars[i].index < uniqueChars[j].index
		})

		// Create a new slice containing only the first members
		result := ""
		for _, pair := range uniqueChars {
			result += string(pair.char)
		}
		if hash == string(result) {
			count++
		}
		if count == 33 {
			fmt.Println(hash)
			break
		}
	}

}
