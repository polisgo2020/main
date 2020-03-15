package main

import (
	"fmt"
)

func main() {
	names := []string{
		"St Petersburg",
		"Moscow",
		"Kazan",
		"Novosibirsk",
		"Vladivostok",
	}

	for i, name := range names {
		// address of the name is not changing
		fmt.Printf("%v %v %s\n", &name, &names[i], name)
	}
}
