package main

import (
	"fmt"
)

func main() {
	fmt.Print("\033[H\033[2J")

	var employee = make(map[string]int)
	employee["Tim"] = 10
	employee["Sandy"] = 20

	// Empty Map
	employeeList := make(map[string]int)

	fmt.Println(len(employee))     // 2
	fmt.Println(len(employeeList)) // 0
	
	// adding to map
	employee["Rocky"] = 30
	employee["Josef"] = 40

	// accessing map
	fmt.Println(employee["Rocky"])

	// updating map
	fmt.Println(employee) // Initial Map
 
	employee["Tim"] = 50 // Edit item
	employee["Josef"] = 10
	employee["Rocky"] = 90
	employee["Sandy"] = 40
	fmt.Println(employee)

	// iterating over map

	for key, element := range employee {
        fmt.Println("Key:", key, "=>", "Element:", element)
    }

	// deleting map
	delete(employee, "Josef")
	delete(employee, "Sandy")
	delete(employee, "Rocky")
	fmt.Println(employee)

	// merging map
	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}
 
	for k, v := range second {
		first[k] = v
	}
 
	fmt.Println(first)

	fmt.Println()
}
