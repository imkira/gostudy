package main

import "fmt"

func showSlice(s []int) {
	fmt.Printf("%#v, len=%d cap=%d\n", s, len(s), cap(s))
}

func main() {
	// start block OMIT
	var a []int = nil
	a = append(a, 1)
	showSlice(a)

	b := make([]int, 0)
	b = append(b, 2)
	showSlice(b)

	c := make([]int, 0, 50)
	c = append(c, 3)
	showSlice(c)

	d := make([]int, 4)
	d = append(d, 4)
	showSlice(d)
	// end block OMIT
}
