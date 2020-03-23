package main

import (
	"fmt"
)

func testPanic() {
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Printf("error: %s\n", err)
	// 	}
	// }()

	fmt.Println("parallel work")
	panic(fmt.Errorf("panic"))
}

func main() {
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Printf("error: %s\n", err)
	// 	}
	// }()

	fmt.Println("step 1")
	go testPanic()
	fmt.Println("step 2")
	fmt.Scanln()
}
