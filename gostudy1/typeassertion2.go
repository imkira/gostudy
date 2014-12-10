package main

import "fmt"

// start block OMIT
func showType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	showType("Hello")
	showType(1)
	showType(true)
}

// end block OMIT
