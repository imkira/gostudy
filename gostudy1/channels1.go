package main

import "fmt"

func main() {
	// start block OMIT
	c := make(chan int, 3)
	c <- 5
	c <- 9
	c <- 10
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	val, ok := <-c
	fmt.Println(val, ok)
	// end block OMIT
}
