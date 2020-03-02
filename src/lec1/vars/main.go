package main

import (
	"fmt"
	"unicode/utf8"
)

const helloStr = "hello"

const (
	constA = iota + 1
	constB
	_
	constC
)

func main() {
	fmt.Println("hello")

	testString := helloStr
	testInt := 123
	var testBool bool

	fmt.Println(testString, testInt, testBool)

	testInt++

	fmt.Println(testString, testInt, testBool)

	// string length
	fmt.Println(len(testString))
	testString = testString + "хелло"
	fmt.Println(len(testString))
	fmt.Println(utf8.RuneCountInString(testString))

	// iota
	fmt.Println(constA, constB, constC)

	// pointers
	ptr := &testString
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = "new"
	fmt.Println(testString)
}
