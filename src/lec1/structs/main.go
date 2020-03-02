package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pet struct {
	name string
}

type user struct {
	name  string
	age   int8
	pets  []pet
	pets2 struct {
		a int
	}
}

func (u *user) AddPet(p pet) {
	u.pets = append(u.pets, p)
}

type owner1 user

type owner2 struct {
	*user
}

func (o2 *owner2) AddPet(p pet) {
	fmt.Println(p)
}

func main() {
	u := user{
		name: "name",
		age:  18,
		// pets: []pet{},
	}

	// u.AddPet(pet{
	// 	name: "name of pet",
	// })

	fmt.Println(u)

	o2 := owner2{
		user: &u,
	}

	o2.AddPet(pet{
		name: "name of pet",
	})

	o2.user.AddPet(pet{
		name: "name of pet",
	})

	fmt.Println(o2.user)

	fmt.Printf("%#v\n", strings.Split("a a a", " "))

}
