package main

import (
	"fmt"
	"os"
)

func Challenge_02() {

	var ops = make(map[byte]func(int) int)
	ops['#'] = func(x int) int { return x + 1 }
	ops['@'] = func(x int) int { return x - 1 }
	ops['*'] = func(x int) int { return x * x }
	ops['&'] = func(x int) int { return x }

	input, err := os.ReadFile("message_02.txt")
	if err != nil {
		panic(err)
	}

	var result string
	var num int = 0
	for _, char := range input {

		if char == '&' {
			result = result + fmt.Sprintf("%d", num)
		}

		num = ops[char](num)

	}
	fmt.Println(result)
}
