package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))

	// with errors
	fmt.Println(sumErr(1, 2, 0))
	fmt.Println(sumErr(1, 1, 0))
	if _, err := sumErr(1, 1, 0); err != nil {
		fmt.Println(err)
	}

	// slice param
	sliceParam([]int{1, 5, 10}...)

	power2 := func(a int) int {
		return a * a
	}

	func(a int) {
		a++
	}(5)

	prefix := "prefix"

	type pwrFunc func(a int) int

	echo := func(a int, lambda pwrFunc) {
		mult := func(a int) string {
			return fmt.Sprintf("%s%d", prefix, a)
		}
		fmt.Println(prefix, mult(lambda(a)))
	}

	echo(1, power2)
	echo(2, power2)
}

func sum(a, b int) (int, int) {
	c := a + b
	return c, c % 2
}

func sumErr(a, b, c int) (int, error) {
	d := a + b
	if d%2 == 1 {
		return 0, errors.New("d % 2 == 1")
	}
	return d, nil
}

func sliceParam(in ...int) {
	for _, i := range in {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
