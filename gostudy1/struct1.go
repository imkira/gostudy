package main

import "fmt"

// start block OMIT

type hoge struct {
	fuga int
	piyo string
}

func (h *hoge) addFuga(x int) {
	h.fuga += x
}

func main() {
	h1 := hoge{fuga: 1, piyo: "Hello, 世界"}
	h1.addFuga(3)
	fmt.Println(h1.fuga)
}

// end block OMIT
