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
	var frame = 1
	var throw = 0
	for index, val := range splitString {

		if frame >= 11 {
			return score
		}

		if val == "X" {
			frame++
			if frame < 10 && (index+2 < len(splitString)) {
				score = score + ConvertStringToValue(val) + ConvertStringToValue(splitString[index+1]) + ConvertStringToValue(splitString[index+2])
				throw = throw + 2
			} else {
				score = score + 10 + ConvertStringToValue(splitString[index+1]) + ConvertStringToValue(splitString[index+2])
				throw++
			}

		} else if val == "/" {

			if frame < 10 && (index+1 < len(splitString)) {
				score = score + ConvertStringToValue(val) + ConvertStringToValue(splitString[index+1])
			} else {
				score = score + 5 + ConvertStringToValue(splitString[index+1])
				frame = 11
			}
			frame++
			throw++
		} else {
			score = score + ConvertStringToValue(val)
			throw++
			if throw%2 == 0 && frame != 10 {
				frame++
			}
		}

	}
	return score
}

func main() {

	fmt.Println(SplitInputString(os.Args[1]))
}
