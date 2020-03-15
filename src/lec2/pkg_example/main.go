package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/polisgo2020/main/src/lec2/pkg_example/processor"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter coefficient")
		return
	}

	c, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Argument is not a number")
		return
	}

	p := processor.NewProcessor(c)
	fmt.Println(p.Number())
}
