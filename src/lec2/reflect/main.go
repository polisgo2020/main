package main

import (
	"fmt"
	"reflect"
)

type table struct {
	Name     string `example:"value1"`
	Category string `example:"value2"`
}

func (t *table) SetName(name string) {
	t.Name = name
}

func main() {
	tbl := table{}
	tbl.SetName("name")

	v := reflect.ValueOf(tbl)
	fmt.Println(v)

	t := reflect.TypeOf(tbl)
	fmt.Println(t)

	if v.Kind() == reflect.Struct {
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
			fmt.Println(t.Field(i).Tag)
			fmt.Println("example tag: ", t.Field(i).Tag.Get("example"))
			fmt.Println()
		}
	}
}
