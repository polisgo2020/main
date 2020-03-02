package main

import (
	"fmt"
)

const Size = 4

func main() {
	strs := [3]string{"a", "b", "c"}
	fmt.Printf("%+v\n", strs)

	// size := int(4)
	// strs_size := [size]string{}
	strsSize := [Size]string{}
	fmt.Printf("%#v\n", strsSize)

	strsSlice := []string{}
	fmt.Printf("%#v len %d cap %d\n", strsSlice, len(strsSlice), cap(strsSlice))

	strsSlice = append(strsSlice, "a", "b", "c")
	fmt.Printf("%#v len %d cap %d\n", strsSlice, len(strsSlice), cap(strsSlice))

	strsSlice = append(strsSlice, "d")
	fmt.Printf("%#v len %d cap %d\n", strsSlice, len(strsSlice), cap(strsSlice))

	strsSlice2 := make([]string, 0, 3)
	fmt.Printf("%#v len %d cap %d\n", strsSlice2, len(strsSlice2), cap(strsSlice2))
	strsSlice2 = append(strsSlice2, strsSlice...)
	fmt.Printf("%#v len %d cap %d\n", strsSlice2, len(strsSlice2), cap(strsSlice2))

	fmt.Println(strsSlice2[0])
	// fmt.Println(strsSlice2[10])
}
