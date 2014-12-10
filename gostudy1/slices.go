package main

import "fmt"

// start block OMIT
func showSlice(s []int) {
	fmt.Printf("%#v, len=%d cap=%d\n", s, len(s), cap(s))
}

func main() {
	var a []int = nil
	showSlice(a)

	b := make([]int, 0)
	showSlice(b)

	c := make([]int, 0, 50)
	showSlice(c)

	d := make([]int, 5)
	showSlice(d)

	e := make([]int, 5, 50)
	showSlice(e)
}

// end block OMIT
