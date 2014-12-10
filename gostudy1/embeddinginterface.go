package main

import "fmt"

// start block OMIT
type number interface{}

type adder interface {
	Add(x, y number) number
}

type multiplier interface {
	Multiply(x, y number) number
}

type calculator interface {
	adder
	multiplier
}

// end block OMIT

// start intCalculator OMIT
type intCalculator struct{}

func (ic *intCalculator) Add(x, y number) number {
	return x.(int) + y.(int)
}

func (ic *intCalculator) Multiply(x, y number) number {
	return x.(int) * y.(int)
}

func main() {
	var calc calculator = &intCalculator{}

	fmt.Println(calc.Add(5, 10))
	fmt.Println(calc.Multiply(5, 10))
}

// end intCalculator OMIT
