package main

import "fmt"

func showMap(m map[string]int) {
	fmt.Printf("%#v, len=%d\n", m, len(m))
}

func main() {
	// start block OMIT
	a := map[string]int{
		"hoge": 5,
		"fuga": 7,
	}
	showMap(a)

	val := a["piyo"]
	fmt.Println(val)

	val, ok := a["piyo"]
	fmt.Println(val, ok)

	val, ok = a["hoge"]
	fmt.Println(val, ok)

	val = a["fuga"]
	fmt.Println(val)
	// end block OMIT
}
