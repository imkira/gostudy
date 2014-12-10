package main

import "fmt"

func main() {
	// start block OMIT
	c := make(chan int, 3)
	c <- 5
	c <- 9
	c <- 10
	for {
		select {
		case x := <-c:
			fmt.Println(x)
		default:
			return
		}
	}
	// end block OMIT
}
