package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func Perm(n int) []int

func main() {
	fmt.Print("\033[H\033[2J")

	// generating slice
	randomNumbers := [4]int{4, 3, 6, 2}
	myslice := randomNumbers[1:3]

	//Provide seed
    rand.Seed(time.Now().Unix())

    //Generate a random array of length n
	temp := rand.Perm(10)
    fmt.Println(append(myslice, temp...))

	fmt.Println()
}
