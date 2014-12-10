package main

import "fmt"

func main() {
	// start block OMIT
	c := make(chan int, 3)
	for {
		select {
		case c <- 1:
			fmt.Println("ok")
		default:
			return
		}
	}
	// end block OMIT
}
