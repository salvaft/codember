package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Challenge_03() {
	var input, err = os.ReadFile("encryption_policies.txt")
	if err != nil {
		panic("File not found")
	}
	var data = strings.Split(string(input), "\n")
	var count int

	for _, key := range data {
		// Split the key into parts
		parts := strings.Split(key, " ")

		// Check if we have enough parts
		if len(parts) != 3 {

			panic(fmt.Sprint("Invalid input ", key))
		}

		// Assign parts to variables
		limitsStr, letter, password := parts[0], string(rune(parts[1][0])), parts[2]
		limits := strings.Split(limitsStr, "-")
		if len(limits) != 2 {
			panic(fmt.Sprint("Invalid input ", key))
		}
		lower, err := strconv.Atoi(limits[0])
		if err != nil {
			panic(fmt.Sprint("Invalid input ", key))
		}
		upper, err := strconv.Atoi(limits[1])
		if err != nil {
			panic(fmt.Sprint("Invalid input ", key))
		}
		if !keyIsValid(password, rune(letter[0]), [2]int{lower, upper}) {
			count++
			if count == 42 {
				fmt.Println(password)
			}
		}

	}

}
func keyIsValid(key string, letter rune, limits [2]int) bool {
	var count int
	for _, char := range key {
		if char == letter {
			count++
		}
	}
	if count < limits[0] || count > limits[1] {
		return false

	}
	return true
}
