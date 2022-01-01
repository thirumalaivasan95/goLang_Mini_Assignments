package main

import "fmt"

func main() {
	fmt.Print("\033[H\033[2J")

	var largerNumber int
	slice := []int{1, 19, 16, 13, 21}

	for _, element := range slice {
		if element > largerNumber {
			largerNumber = element
		}
	}
	fmt.Println("Largest number of slice is:", largerNumber)

	fmt.Printf("\n")
}
