package main

import "fmt"

// start block OMIT
func last(first int, rest ...int) int {
	if len(rest) == 0 {
		return first
	}
	return rest[len(rest)-1]
}

func main() {
	fmt.Println(last(1))
	fmt.Println(last(4, 5))
	fmt.Println(last(6, 7, 8))
}

// end block OMIT
