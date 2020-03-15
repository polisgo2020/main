package main

import (
	"fmt"
)

type iFurniture interface {
	SetName(in string)
}

type iTable interface {
	GetName() string
	iFurniture
}

type table struct {
	name string
}

func (t table) GetName() string {
	return t.name
}

func (t *table) SetName(in string) {
	t.name = in
}

func main() {
	t1(table{})
	t2(&table{}) // pointer is needed because SetName has method with pointer t *table
}

func t1(something interface{}) {
	switch v := something.(type) {
	case iTable:
		fmt.Println(v.GetName())
	}
}

func t2(t iTable) {
	fmt.Println(t.GetName())
}

const nil = 123
