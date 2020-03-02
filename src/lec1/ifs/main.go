package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("hw")
	}
	if 1 == 1 {
		fmt.Println("strong typing")
	} else if 1 == 2 {
		fmt.Println("multiple if")
	}

	strs := map[string]int{
		"a": 1,
	}
	if el, ok := strs["a"]; ok {
		fmt.Println(el)
	}
}
