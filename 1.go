package main

import "fmt"

func main() {

	fmt.Print("\033[H\033[2J")

	var i, j int

	for i = 4; i >= 0; i-- {
		for j = i; j >= 0; j-- {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	fmt.Println()
}
