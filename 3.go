package main

import (
	"fmt"
)

func switchCase(s string) {
	fmt.Printf("%s\n\n", s)

	var vowel, consonant, digits, specialCharacter int

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U', 'y', 'Y':
			vowel += 1

		case 'b', 'B', 'c', 'C', 'd', 'D', 'f', 'F', 'g', 'G', 'h', 'H',
			'j', 'J', 'k', 'K', 'l', 'L', 'm', 'M', 'n', 'N', 'p', 'P',
			'q', 'Q', 'r', 'R', 's', 'S', 't', 'T', 'v', 'V', 'w', 'W',
			'x', 'X', 'z', 'Z':
			consonant += 1

		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			digits += 1

		default:
			specialCharacter += 1
		}
	}
	fmt.Printf("Vowels: %d\n", vowel)
	fmt.Printf("Consonant: %d\n", consonant)
	fmt.Printf("Digits: %d\n", digits)
	fmt.Printf("Special Characters: %d\n", specialCharacter)
}

func main() {
	fmt.Print("\033[H\033[2J")

	var str string = "Quest-global company1111"
	switchCase(str)

	fmt.Println()
}
