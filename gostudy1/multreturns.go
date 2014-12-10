package main

import "fmt"

// start block OMIT
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(swap(1, 2))
}

// end block OMIT
