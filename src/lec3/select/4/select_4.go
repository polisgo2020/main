package main

import (
	"fmt"
)

func main() {
	in := make(chan int, 0)

	go func(out chan<- int) {
		for i := 0; i <= 3; i++ {
			fmt.Println("before", i)
			out <- i
			fmt.Println("after", i)
		}
		close(out)
		fmt.Println("generator finish")
	}(in)

Loop1:
	for {
		select {
		case i, ok := <-in:
			if !ok {
				break Loop1
			}
			fmt.Println("\tget", i)
		}
	}
}
