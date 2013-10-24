package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SplitInputString(input string) []string {
	return strings.Split(input, "-")
}

func ConvertStringToValue(input string) int {
	if input == "X" {
		return 10
	}

	if input == "/" {
		return 5
	}

	val, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return 0
	}

	return int(val)
}

func BasicScore(input string) int {
	var splitString = SplitInputString(input)

	var score int = 0
	for _, val := range splitString {
		score = score + ConvertStringToValue(val)
	}
	return score
}

func main() {

	fmt.Println(SplitInputString(os.Args[1]))
}
