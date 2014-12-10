package main

import "fmt"

// start block OMIT

type hoge struct {
	fuga // Embedding

	piyo string
}

type fuga struct {
	num int
}

func (f *fuga) add(x int) {
	f.num += x
}

func main() {
	h1 := hoge{fuga: fuga{num: 1}, piyo: "Hello, 世界"}
	h1.add(3)
	fmt.Println(h1.num)
	h1.fuga.add(2)
	fmt.Println(h1.fuga.num)
}

// end block OMIT
