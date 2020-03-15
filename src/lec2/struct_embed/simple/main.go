package main

import (
	"fmt"
)

type furniture struct {
	name  string
	price int
}

type table struct {
	length int
	name   string
	furniture
}

func main() {
	t := table{
		length: 10,
		furniture: furniture{
			name:  "name",
			price: 200,
		},
	}

	fmt.Printf("%#v\n", t)
	fmt.Println(t.name)
	fmt.Println(t.furniture.name)
	fmt.Println(t.price)
}
