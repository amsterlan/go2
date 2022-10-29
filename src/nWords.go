package main

import (
	"fmt"
)

type words []rune

func numbWord(c rune) bool {
	nWords := words{' '}
	for _, value := range nWords {
		if value == c {
			return true
		}
	}
	return false
}

func main() {
	word := "um dois tres quatro"
	v := 1
	for _, value := range word {
		if numbWord(value) {
			v++
		}
	}
	fmt.Printf("%d Words", v)
}
