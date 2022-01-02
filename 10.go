// This sample program demonstrates how to create goroutines and
// how the scheduler behaves.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// main is the entry point for all Go programs.
func main() {
	fmt.Print("\033[H\033[2J")
	// Allocate 1 logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
	// wg is used to wait for the program to finish.
	// Add a count of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines")
	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()
		// Display the alphabet three times
		for count := 0; count < 1; count++ {
			for char := 0; char <= 100; char++ {
				fmt.Printf("%d ", char)
			}
		}
		fmt.Println()
		fmt.Println()
	}()
	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()
		// Display the alphabet three times
		for count := 0; count < 1; count++ {
			for char := 0; char <= 100; char++ {
				fmt.Printf("%d ", char)
			}
		}
		fmt.Println()
		fmt.Println()
	}()
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
	fmt.Println()
}
