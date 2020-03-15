package main

import (
	"fmt"
)

func main() {
	var in string
	// UTF string
	in = "Привет"
	// iterating over []rune
	for _, c := range in {
		fmt.Print(string(c))
	}
	fmt.Println()
	// iterating over []byte
	for i := 0; i < len(in); i++ {
		fmt.Print(string(in[i]))
	}
	fmt.Println()
	// iterating over []rune
	runeIn := []rune(in)
	for i := 0; i < len(runeIn); i++ {
		fmt.Print(string(runeIn[i]))
	}
	fmt.Println()
}
