package main

import (
	"fmt"
	"time"
)

// start block OMIT

func writer(c chan int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	go writer(c)

	for x := range c {
		fmt.Println(x)
	}
}

// end block OMIT
