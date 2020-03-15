package main

import (
	"fmt"
	"unsafe"
)

type Example struct {
	BoolValue   bool
	IntValue    int16
	FloatValue  float32
	FloatValue2 float32
}

func main() {
	example1()
}

func example1() {
	firstval := 55555555.55

	someval := 555555555599

	example := &Example{
		BoolValue:   true,
		IntValue:    10,
		FloatValue:  3.141592,
		FloatValue2: 3.141592,
	}

	exampleNext := &Example{
		BoolValue:   true,
		IntValue:    10,
		FloatValue:  3.141592,
		FloatValue2: 3.141592,
	}

	fmt.Printf("%T size: %d, addr: %v (%d)\n", firstval, unsafe.Sizeof(firstval), &firstval, &firstval)

	fmt.Printf("%T size: %d, addr: %v (%d)\n", someval, unsafe.Sizeof(someval), &someval, &someval)

	printInfo(example)

	fmt.Println("NEXT")

	printInfo(exampleNext)
}

func printInfo(example *Example) {
	alignmentBoundary := unsafe.Alignof(example)

	sizeBool := unsafe.Sizeof(example.BoolValue)
	offsetBool := unsafe.Offsetof(example.BoolValue)

	sizeInt := unsafe.Sizeof(example.IntValue)
	offsetInt := unsafe.Offsetof(example.IntValue)

	sizeFloat := unsafe.Sizeof(example.FloatValue)
	offsetFloat := unsafe.Offsetof(example.FloatValue)

	sizeFloat2 := unsafe.Sizeof(example.FloatValue2)
	offsetFloat2 := unsafe.Offsetof(example.FloatValue2)

	fmt.Printf("Alignment Boundary: %d\n", alignmentBoundary)

	fmt.Printf("%T Value = Size: %d Offset: %d Addr: %v (%d)\n",
		example.BoolValue, sizeBool, offsetBool, &example.BoolValue, &example.BoolValue)

	fmt.Printf("%T Value = Size: %d Offset: %d Addr: %v (%d)\n",
		example.IntValue, sizeInt, offsetInt, &example.IntValue, &example.IntValue)

	fmt.Printf("%T Value = Size: %d Offset: %d Addr: %v (%d)\n",
		example.FloatValue, sizeFloat, offsetFloat, &example.FloatValue, &example.FloatValue)

	fmt.Printf("%T Value = Size: %d Offset: %d Addr: %v (%d)\n",
		example.FloatValue2, sizeFloat2, offsetFloat2, &example.FloatValue2, &example.FloatValue2)
}
