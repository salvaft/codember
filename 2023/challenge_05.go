package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func Challenge_05() {
	input, err := os.ReadFile("database_attacked.txt")
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}
	lines := strings.Split(string(input), "\n")

	var result string
	for _, line := range lines {
		values := strings.Split(line, ",")

		id := values[0]
		username := values[1]
		email := values[2]
		fmt.Println(id, username, email)
		if !isAlphanumeric(id) || !isValidEmail(email) || !isAlphanumeric(username) {
			result += string(username[0])
			continue
		}

	}
	fmt.Println(result)
}

func isAlphanumeric(str string) bool {
	if str == "" {
		return false
	}
	for _, char := range str {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}

	}
	return true
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	// Define a regular expression for a simple email validation
	// This example doesn't cover all possible cases; adjust it based on your requirements.
	// For a more comprehensive email validation, consider using a specialized library.

	return emailRegex.MatchString(email)
}
