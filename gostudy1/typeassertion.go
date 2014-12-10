package main

import "fmt"

func main() {
	// start block OMIT
	var x interface{}

	x = "Hello"
	s := x.(string)
	fmt.Println(s + " 世界")

	// fail
	i, ok := x.(int)
	fmt.Println(i, ok)

	// panic
	i = x.(int)
	fmt.Println(i)
	// end block OMIT
}
