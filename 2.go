package main

import "fmt"
  
func main() {

	fmt.Print("\033[H\033[2J")

    day := "Tue"

    switch {
    case day == "Mon":
        fmt.Println("Monday")
        fallthrough
    case day == "Tue":
        fmt.Println("Tuesday")
        fallthrough
    case day == "Wed":
        fmt.Println("Wednesday")
		fallthrough
	case day == "Thur":
        fmt.Println("Thursday")
		fallthrough
	case day == "Fri":
        fmt.Println("Friday")
    }
}