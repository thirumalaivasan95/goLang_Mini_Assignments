package main

import "fmt"

func main() {
	fmt.Print("\033[H\033[2J")

	//Initialize an array
	inputArray := []int{10, 20, 30, 56, 67, 90, 10, 20}
	printUniqueValue(inputArray)

	fmt.Println()
}

func printUniqueValue(arr []int) {
	//Create a dictionary of values for each element using map
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	fmt.Println(dict)
}
