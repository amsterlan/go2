package main

import (
	"fmt"
)

type chars []rune

func vowels(c rune) bool {
	vowels := chars{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, value := range vowels {
		if value == c {
			return true
		}
	}
	return false
}

func main() {
	word := "palavra"
	v := 0
	for _, value := range word {
		if vowels(value) {
			v++
		}
	}
	fmt.Printf("%d Vowels", v)
}
