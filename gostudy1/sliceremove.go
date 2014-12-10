package main

import "fmt"

func showSlice(s []int) {
	fmt.Printf("%#v, len=%d cap=%d\n", s, len(s), cap(s))
}

func main() {
	// start block OMIT
	a := []int{4, 8, 15, 16, 23, 42}
	showSlice(a)

	// remove first
	b := a[1:]
	showSlice(b)

	// remove last
	b = a[0 : len(a)-1]
	showSlice(b)

	// remove second
	b = append(a[:2], a[3:]...)
	showSlice(b)

	// changed
	showSlice(a)

	// end block OMIT
}
