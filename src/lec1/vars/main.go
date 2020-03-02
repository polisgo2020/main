package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	testString := "hello"
	testInt := 123
	var testBool bool

	fmt.Println(testString, testInt, testBool)

	testInt++

	fmt.Println(testString, testInt, testBool)
}
