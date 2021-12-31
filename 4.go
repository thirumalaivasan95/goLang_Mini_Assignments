package main

import (
	"fmt"
)

func reverse(s string) string {
	inputString := []rune(s)
	for i, j := 0, len(inputString)-1; i < j; i, j = i+1, j-1 {
		inputString[i], inputString[j] = inputString[j], inputString[i]
	}
	return string(inputString)
}

func main() {
	fmt.Print("\033[H\033[2J")

	str := "Quest Global"
	reverseString := reverse(str)
	fmt.Println(str)
	fmt.Println(reverseString)

	fmt.Println()
}
