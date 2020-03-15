package main

import (
	"fmt"
)

func main() {
	stringsMap := map[string]int{}
	stringsMap["a"] = 1
	stringsMap["b"] = 2
	fmt.Printf("%#v\n", stringsMap)
	fmt.Printf("%#v\n", stringsMap["a"])
	fmt.Printf("%#v\n", stringsMap["d"])

	x := X{}
	x.add("a", 1)
	x.add("b", 2)
	fmt.Printf("%#v\n", x)
}

type X map[string]int

func (x *X) add(k string, v int) {
	(*x)[k] = v
}

func (x X) addNoPtr(k string, v int) {
	x[k] = v
}
