package main

import (
	"fmt"
)

func main() {
	strs := [3]string{"a", "b", "c"}
	fmt.Printf("%+v\n", strs)
	fmt.Println(strs[1])

	const Size = 4
	// size := int(4)
	// strs_size := [size]string{}
	strsSize := [Size]string{}
	fmt.Printf("%#v\n", strsSize)

	// slices
	strsSlice := []string{}
	fmt.Printf("%#v len %d cap %d\n", strsSlice, len(strsSlice), cap(strsSlice))

	strsSlice = append(strsSlice, "a", "b", "c")
	fmt.Printf("%#v len %d cap %d\n", strsSlice, len(strsSlice), cap(strsSlice))

	strsSlice = append(strsSlice, "d")
	fmt.Printf("%#v len %d cap %d\n", strsSlice, len(strsSlice), cap(strsSlice))

	// make
	strsSlice2 := make([]string, 0, 3)
	fmt.Printf("%#v len %d cap %d\n", strsSlice2, len(strsSlice2), cap(strsSlice2))
	strsSlice2 = append(strsSlice2, strsSlice...)
	fmt.Printf("%#v len %d cap %d\n", strsSlice2, len(strsSlice2), cap(strsSlice2))

	fmt.Println(strsSlice2[0])
	// fmt.Println(strsSlice2[10])

	// slicing slices
	// take [1,3) elements by making the link
	strsSlice3 := strsSlice2[1:3]
	fmt.Println(strsSlice3)

	// copy slices
	sliceCopy := strsSlice2[:]
	fmt.Println(sliceCopy)

	strsSlice2[0] = "X"
	fmt.Println(strsSlice2, sliceCopy)

	sliceCopy = append(sliceCopy, "x", "y", "z")
	fmt.Println(strsSlice2, sliceCopy)

	// correct copy
	sliceCopy2 := make([]string, len(sliceCopy), len(sliceCopy))
	copy(sliceCopy2, sliceCopy)
	fmt.Println(sliceCopy, sliceCopy2)
}
