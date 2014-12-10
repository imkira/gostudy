package main

import "fmt"

// start block OMIT
func showMap(m map[string]int) {
	fmt.Printf("%#v, len=%d\n", m, len(m))
}

func main() {
	var a map[string]int = nil
	showMap(a)

	b := make(map[string]int, 0)
	showMap(b)

	c := make(map[string]int, 5)
	showMap(c)

	d := map[string]int{
		"hello": 1,
		"world": 2,
	}
	showMap(d)
}

// end block OMIT
