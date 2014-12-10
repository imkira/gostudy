package main

import "fmt"

// start block OMIT
func fibonacci() func() int {
	i, j := 0, 1
	return func() int {
		k := i
		i, j = j, j+i
		return k
	}
}

func main() {
	nextFibonacci := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(nextFibonacci())
	}
}

// end block OMIT
