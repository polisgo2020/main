package main

import (
	"testing"
)

func TestBubble(t *testing.T) {
	in := []int{0, -500, 1, 2, -1}
	expected := []int{-500, -1, 0, 1, 2}
	actual := bubblesort(in)
	if !testEq(expected, actual) {
		t.Errorf("%v is not eqal to expected %v", actual, expected)
	}
}

func TestMerge(t *testing.T) {
	in := []int{0, -500, 1, 2, -1}
	expected := []int{-500, -1, 0, 1, 2}
	actual := mergesort(in)
	if !testEq(expected, actual) {
		t.Errorf("%v is not eqal to expected %v", actual, expected)
	}
}

func testEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

/*
	go test -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html

*/
