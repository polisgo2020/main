package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	condition := true
	var i int32
	for condition {
		fmt.Print(i)
		i++
		condition = i < 10
	}
	fmt.Println()

	i = 0
	for {
		fmt.Print(i)
		i++
		if i >= 10 {
			break
		}
	}
	fmt.Println()

	i = 0
forCycle:
	for {
		switch i % 2 {
		case 0:
			i++
			continue
		case 1:
			fmt.Print(i)
		}

		i++

		switch i {
		case 10:
			break forCycle
		}
	}
	fmt.Println()
}
