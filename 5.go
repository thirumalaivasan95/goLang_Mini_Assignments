package main

import (
	"fmt"
)

func main() {
	fmt.Print("\033[H\033[2J")

	arr2D := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(arr2D)

	// without range:

	row := len(arr2D)
	column := len(arr2D[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Println(arr2D[i][j])
		}
	}
	
	fmt.Printf("\n\n")
	// with range:

	for _, row := range arr2D {
        for _, val := range row {
            fmt.Println(val)
        }
    }

	fmt.Printf("\n")

	fmt.Println()
}