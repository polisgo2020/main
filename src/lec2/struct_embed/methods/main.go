package main

import (
	"fmt"
)

type furniture struct {
	name  string
	price int
}

func (f *furniture) setName(name string) {
	f.name = name
}

type table struct {
	length int
	name   string
	furniture
}

func (t *table) setName(name string) {
	t.name = name
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
	t.setName("table")
	fmt.Printf("%#v\n", t)
	t.furniture.setName("furniture")
	fmt.Printf("%#v\n", t)
}
