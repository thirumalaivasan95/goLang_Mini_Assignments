package main

import "fmt"

type RuneTypeCounts struct {
	vowels, consonants, digits, specialCharacter int
}

func analyseText(text string) RuneTypeCounts {
	var result RuneTypeCounts
	m := make(map[rune]*int)
	for _, r := range "aeiou" {
		m[r] = &result.vowels
	}
	for _, r := range "bcdfghjklmnpqrstvwxyz" {
		m[r] = &result.consonants
	}
	for _, r := range "0123456789" {
		m[r] = &result.digits
	}
	for _, r := range text {
		ref, inMap := m[r]
		if inMap {
			*ref++
		} else {
			result.specialCharacter++
		}
	}
	return result
}

func main() {
	fmt.Print("\033[H\033[2J")
	rt := analyseText("Quest-global company1111")
	fmt.Printf("%+v", rt)
	fmt.Println()
	fmt.Println()
}
