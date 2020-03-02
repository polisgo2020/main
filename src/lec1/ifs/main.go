package main

import (
	"fmt"
)

func main() {
	if (true || true) && true {
		fmt.Println("hw")
	}
	if 1 == 1 {
		fmt.Println("strong typing")
	} else if 1 == 2 {
		fmt.Println("multiple if")
	} else {
		fmt.Println("else")
	}

	strs := map[string]int{
		"a": 1,
	}
	if el, ok := strs["a"]; ok {
		fmt.Println(el)
	}

	switch len(strs) {
	case 0, 1:
		fmt.Println("len 0 || 1")
	default:
		fmt.Println("len > 1")
	}

	switch {
	case 1 == 2:
		fmt.Println("2 != 1")
	case 1 == 1:
		fmt.Println("1 = 1")
		// fallthrough
	case 2 == 2:
		fmt.Println("2 = 2")
	}

}
