package main

import "fmt"

func showSlice(s []int) {
	fmt.Printf("%#v, len=%d cap=%d\n", s, len(s), cap(s))
}

func main() {
	// start block OMIT
	a := []int{4, 8, 15, 16, 23}
	// 1st
	showSlice(a[0:1])
	// 3rd and 4th
	showSlice(a[2:4])
	// first 4
	showSlice(a[0:4])
	// 2nd to last
	showSlice(a[1:])
	// first 4
	showSlice(a[:4])
	// first 5
	showSlice(a[:5])
	// all
	showSlice(a[:])
	// end block OMIT
}
