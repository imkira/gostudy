package main

import "fmt"

func main() {
	// start block OMIT
	a := []int{4, 8, 15, 16, 23, 42}
	for i, val := range a {
		fmt.Println(i, val)
	}

	for _, val := range a {
		fmt.Println(val)
	}
	// end block OMIT
}
