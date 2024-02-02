package main

import (
	"os"
	"strconv"
)

func main() {
	var challenges = []func(){
		Challenge_01,
		Challenge_02,
		Challenge_03,
		Challenge_04,
		Challenge_05,
	}
	challenge_idx, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	challenges[challenge_idx-1]()
}
