package main

import "fmt"

type cat struct {
}

type dog struct {
}

// start block OMIT
func (c *cat) Sound() string {
	return "ニャン"
}

func (d *dog) Sound() string {
	return "ワンワン"
}

type animal interface {
	Sound() string
}

func play(a animal) {
	fmt.Println(a.Sound())
}

func main() {
	play(&cat{})
	play(&dog{})
}

// end block OMIT
